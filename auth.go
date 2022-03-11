package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/benchttp/server/jwt"
)

func (s *Server) handleSignin(w http.ResponseWriter, r *http.Request) {
	// TODO sync with front-end regarding the data structure
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
	}

	name, email, err := s.OAuthClient.GetUser(ghToken)
	if err != nil {
		writeError(w, &ErrInternal)
	}

	// if user does not exists -> create new user

	// webToken authenticates the user from the webapp.
	webToken, err := createToken(name, email)
	if err != nil {
		writeError(w, &ErrInternal)
	}

	// accessToken authenticates the user from the runner.
	accessToken, err := createToken(name, email)
	if err != nil {
		writeError(w, &ErrInternal)
	}

	writeJSON(w, struct {
		WebToken    string `json:"jwt"`
		AccessToken string `json:"accessToken"`
	}{
		WebToken:    webToken,
		AccessToken: accessToken,
	}, 201)
}

func (s *Server) handleCreateAccessToken(w http.ResponseWriter, r *http.Request) {
	// TODO
	name, email := "marcel patulacci", "marcelpatulacci@policenationale.fr"

	accessToken, err := createToken(name, email)
	if err != nil {
		writeError(w, &ErrInternal)
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
