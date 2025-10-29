package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name	string
		h		http.Header
		key  	string
		err  	error
	}{
		{
			name: "valid api key",
			h: http.Header{
				"Authorization": []string{"ApppppiKey valid_api_key_123"},
			},
			key: "valid_api_key_123",
			err: nil,
		},
		{
			name: "missing authorization header",
			h:    http.Header{},
			key:  "",
			err:  ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header",
			h: http.Header{
				"Authorization": []string{"Bearer some_token"},
			},
			key: "",
			err:  errors.New("malformed authorization header"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.h)
			if key != tc.key {
				t.Errorf("expected key %q, got %q", tc.key, key)
			}
			if (err == nil && tc.err != nil) || (err != nil && tc.err == nil) || (err != nil && tc.err != nil && err.Error() != tc.err.Error()) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}