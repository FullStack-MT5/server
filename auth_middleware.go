package server

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/benchttp/server/benchttp"
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
func (s *Server) mustAuth(hf http.HandlerFunc) http.HandlerFunc {
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

		user, err := s.UserService.GetByCred(claims.Name, claims.Email)
		if err != nil {
			writeError(w, &ErrUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, &user)
		hf(w, r.WithContext(ctx))
	}
}

// userFromContext returns the user that was set on request.Context
// by mustAuth middleware. Returns nil if no user is set in the context.
// A nil User must be treated as an internal error, as userFromContext
// must be called where we expect user to exist.
func userFromContext(ctx context.Context) *benchttp.User {
	if u := ctx.Value(userKey); u != nil {
		return u.(*benchttp.User)
	}
	return nil
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
