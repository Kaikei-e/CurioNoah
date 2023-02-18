package feedCollection

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCollectAll(t *testing.T) {
	// TODO
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v2/fetch-feed/stored-all", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	cases := map[string]struct {
		input   echo.Context
		wantErr bool
	}{
		"success": {
			input:   c,
			wantErr: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := CollectAll(c.input)
			if (err != nil) != c.wantErr {
				t.Errorf("CollectAll() error = %v, wantErr %v", err, c.wantErr)
				return
			}
		})
	}

}
