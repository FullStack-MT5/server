package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	exchangeURL = "https://github.com/login/oauth/access_token"
	userURL     = "https://api.github.com/user"
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
		return "", fmt.Errorf("%w: %s", errRequest, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("%w: %s: %s", errStatusCode, res.Status, err)
	}

	var t struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return "", fmt.Errorf("%s: %s", errParse, err)
	}

	return t.AccessToken, nil
}

func (c OAuthClient) GetUser(token string) (name, email string, err error) {
	req, err := http.NewRequest("GET", userURL, nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", errRequest, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", "", fmt.Errorf("%w: %s: %s", errStatusCode, res.Status, err)
	}

	var u struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		// All the available properties are defined in GitHub docs:
		// https://docs.github.com/en/rest/reference/users#get-the-authenticated-user
	}

	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return "", "", fmt.Errorf("%s: %s", errParse, err)
	}

	return u.Name, u.Email, nil
}
