package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authentication info found")
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) != 2 {
		return "", errors.New("malformed auth header")
	}

	if splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return splitAuth[1], nil
}
