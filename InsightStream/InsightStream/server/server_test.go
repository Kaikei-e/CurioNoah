package server

import (
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	const err404 = "{\n  \"message\": \"Not Found\"\n}"

	reqSuccessAlive := httptest.NewRequest("GET", "/", nil)
	reqSuccessFetchFeeds := httptest.NewRequest("GET", "/api/v1/fetch-baseFeeds", nil)
	reqFailureBase := httptest.NewRequest("GET", "/wahfohofeofegf", nil)
	reqFailureFetchFeeds := httptest.NewRequest("GET", "/api/v1/fetch-feed", nil)

	tests := map[string]struct {
		input      *http.Request
		statusCode int
		wantErr    bool
	}{
		"successAlive": {

			input:      reqSuccessAlive,
			statusCode: 200,
			wantErr:    false,
		},
		"successFetchFeeds": {
			input:      reqSuccessFetchFeeds,
			statusCode: 200,
			wantErr:    false,
		},
		"failure": {
			// Wrong path
			input:      reqFailureBase,
			statusCode: 200,
			wantErr:    true,
		},
		"failureFetchFeeds": {
			// Wrong path
			input:      reqFailureFetchFeeds,
			statusCode: 200,
			wantErr:    true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if cmp.Equal(string(body), err404) {
				t.Errorf("Server() got = %v, want %v", string(body), err404)
				return
			}

			if resp.StatusCode != tt.statusCode {
				t.Errorf("Server() got = %v, want %v.", resp.StatusCode, tt.statusCode)

			}

		})
	}
}
