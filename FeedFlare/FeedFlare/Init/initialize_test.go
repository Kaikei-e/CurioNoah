package Init

import (
	"errors"
	"testing"
)

func TestInitialize(t *testing.T) {
	cases := map[string]struct {
		want    error
		wantErr bool
	}{
		"success": {
			want:    nil,
			wantErr: false,
		},
		"failureGetWorkingDirectory": {
			want:    errors.New("failed to get working directory"),
			wantErr: false,
		},
		"failureEnvNotExist": {
			want:    errors.New("failed to open .env: no such file or directory"),
			wantErr: false,
		},
		"failureEnvNotReadable": {
			want:    errors.New("failed to read .env"),
			wantErr: false,
		},
		"failureEnvNotParsable": {
			want:    errors.New("failed to parse .env"),
			wantErr: false,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			err := Initialize()
			if err != tt.want && tt.wantErr {
				t.Errorf("got %v, want %v", err, tt.want)
			}

		})
	}
}
