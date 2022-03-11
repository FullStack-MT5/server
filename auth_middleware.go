package server

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/benchttp/server/jwt"
)

type contextKey string

const (
	userKey contextKey = "user"
)

// mustAuth is a authentication middleware. It looks for
// a JWT inside the request headers and validates it.
// It the token is valid, mustAuth retreives the user
// associated in the claims and attaches it, if found,
// on request.Context.
func mustAuth(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		signed, err := bearerToken(r)
		if err != nil {
			writeError(w, &ErrUnauthorized)
			return
		}

		token, err := jwt.Verify(signed)
		if err != nil {
			writeError(w, &ErrUnauthorized)
			return
		}

		claims, err := jwt.ClaimsOf(*token)
		if err != nil {
			writeError(w, &ErrUnauthorized)
			return
		}

		// TODO user exists? get user
		user := claims

		ctx := context.WithValue(r.Context(), userKey, &user)
		hf(w, r.WithContext(ctx))
	}
}

// bearerScheme is the string prefixing the key or token
// in the authorization headers: "Bearer <key>".
const bearerScheme = "Bearer "

func bearerToken(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if !strings.HasPrefix(ah, bearerScheme) {
		return "", errors.New("invalid authorization headers")
	}

	return strings.TrimPrefix(ah, bearerScheme), nil
}
