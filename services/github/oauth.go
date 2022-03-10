package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	exchangeURL = "https://github.com/login/oauth/access_token"
)

type OAuthClient struct {
	id     string
	secret string
}

func NewOauth(id, secret string) OAuthClient {
	return OAuthClient{
		id:     id,
		secret: secret,
	}
}

func (c OAuthClient) ExchangeForAccessToken(code string) (string, error) {
	req, err := http.NewRequest("POST", exchangeURL, nil)
	if err != nil {
		return "", err
	}

	// Parameters are expected to be inside the query.
	q := req.URL.Query()
	q.Set("client_id", c.id)
	q.Set("client_secret", c.secret)
	q.Set("code", code)
	req.URL.RawQuery = q.Encode()

	// GitHub defaults to encode the response in plain text.
	// Explicitly asks for JSON encoding.
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// TODO handle 403?

	var t struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	return t.AccessToken, nil
}
