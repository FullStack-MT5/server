package jwt

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidClaims = errors.New("invalid token claims")
	ErrInvalidToken  = errors.New("invalid token")
)

// Claims is the custom claims used to signed JSON web token.
type Claims struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}

// NewClaims returns Claims to provide to Sign.
func NewClaims(name, email string, exp time.Time) Claims {
	return Claims{
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}
}

// Sign returns a new token as a signed a string.
func Sign(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(key())
	if err != nil {
		return "", ErrInvalidClaims
	}

	return signed, nil
}

// Verify parses and validates signed and returns a valid token.
// The token can then be used to access the Claims it holds.
func Verify(signed string) (*jwt.Token, error) {
	// Parse returns an error if the token is invalid
	// or if the signature does not match.
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return key(), nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidClaims
	}

	return token, nil
}

// ClaimsOf returns the underlying claims from token.
func ClaimsOf(token jwt.Token) (Claims, error) {
	m, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Claims{}, ErrInvalidClaims
	}

	claims := Claims{}

	err := decodeMapClaims(m, &claims)
	if err != nil {
		return Claims{}, err
	}

	return claims, nil
}

func decodeMapClaims(m jwt.MapClaims, dst interface{}) error {
	e, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(e, dst); err != nil {
		return err
	}

	return nil
}
