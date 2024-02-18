package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("no auth header", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})

	t.Run("malformed auth header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer")
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
