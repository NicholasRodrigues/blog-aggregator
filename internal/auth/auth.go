package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("not auth header included in request")

// GetBearerToken header format: Authorization
func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

// GetApiKey header format: Authorization: ApiKey <key>
func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}

	splitAuth := strings.Split(authHeader, " ")

	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
