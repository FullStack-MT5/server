package server

import (
	"encoding/json"
	"net/http"
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

	token, err := s.OAuthClient.ExchangeForAccessToken(body.Code)
	if err != nil {
		writeError(w, &ErrBadRequest)
	}

	name, email, err := s.OAuthClient.GetUser(token)
	if err != nil {
		writeError(w, &ErrInternal)
	}

	// if user does not exists -> create new user
	// sign jwt
	// send back jwt
}
