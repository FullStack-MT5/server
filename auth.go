package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/benchttp/server/benchttp"
	"github.com/benchttp/server/jwt"
)

func (s *Server) handleSignin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Code string `json:"code"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeError(w, &ErrBadRequest)
		return
	}

	ghToken, err := s.OAuthClient.ExchangeForAccessToken(body.Code)
	if err != nil {
		writeError(w, &ErrBadRequest)
		return
	}

	name, email, err := s.OAuthClient.GetUser(ghToken)
	if err != nil {
		writeError(w, &ErrInternal)
		return
	}

	user := benchttp.User{}

	switch s.UserService.Exists(email) {
	case true:
		user, err = s.UserService.GetByEmail(email)
	default:
		user, err = s.UserService.Create(name, email)
	}

	if err != nil {
		writeError(w, &ErrUnauthorized)
		return
	}

	// webToken authenticates the user from the webapp.
	webToken, err := createToken(user.Name, user.Email)
	if err != nil {
		writeError(w, &ErrInternal)
		return
	}

	writeJSON(w, struct {
		WebToken string `json:"jwt"`
	}{
		WebToken: webToken,
	}, 201)
}

func (s *Server) handleCreateAccessToken(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r.Context())
	if user == nil {
		writeError(w, &ErrInternal)
		return
	}

	accessToken, err := createToken(user.Name, user.Email)
	if err != nil {
		writeError(w, &ErrInternal)
		return
	}

	writeJSON(w, struct {
		AccessToken string `json:"accessToken"`
	}{
		AccessToken: accessToken,
	}, 201)
}

func createToken(name, email string) (string, error) {
	claims := jwt.NewClaims(name, email, time.Now().Add(24*time.Hour))
	return jwt.Sign(claims)
}
