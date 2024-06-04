package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the request headers
// Example:
// Authorization: APIKey {insert api_key here}
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(authHeader, " ")
	if len(vals) != 2 || vals[0] != "APIKey" {
		return "", errors.New("malformed auth header")
	}
	return vals[1], nil
}
