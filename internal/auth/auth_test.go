package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		wantKey string
		wantErr error
	}{
		{
			name:    "no auth header",
			header:  http.Header{},
			wantKey: "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "improper formatting of header",
			header: http.Header{
				"Authorization": []string{"Bearer xyz"},
			},
			wantKey: "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "fine",
			header: http.Header{
				"Authorization": []string{"ApiKey 12345"},
			},
			wantKey: "12345",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)

			if got != tt.wantKey {
				t.Errorf("GetAPIKey() got key = %v, want %v", got, tt.wantKey)
			}

			// Sprawdzenie błędów po stringu, bo errors.New tworzy nowy obiekt
			if (err != nil && tt.wantErr == nil) ||
				(err == nil && tt.wantErr != nil) ||
				(err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
