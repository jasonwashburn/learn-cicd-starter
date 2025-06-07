package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "no auth header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "no api key",
			headers: http.Header{
				"Authorization": []string{"Bearer 1234567890"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234567890"},
			},
			want:    "1234567890",
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetAPIKey(test.headers)
			if (err != nil) != test.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, test.want)
			}
		})
	}
}
