package adaptor

// import (
// 	"errors"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/labstack/echo/v4"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestCollectAll(t *testing.T) {

// 	// TODO implement various test cases
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/api/v2/fetch-feed/stored-all", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	cases := map[string]struct {
// 		input   echo.Context
// 		wantErr bool
// 	}{
// 		"success": {
// 			input:   c,
// 			wantErr: false,
// 		},
// 		"failure": {
// 			input:   nil,
// 			wantErr: true,
// 		},
// 	}

// 	for name, c := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			err := CollectAll(c.input, cl)

// 			if errors.Is(err, errors.New("context is nil")) {
// 				t.Errorf("CollectAll() got = %v, want %v", err, errors.New("context is nil"))
// 				return

// 			}

// 			if (err != nil) != c.wantErr {
// 				t.Errorf("CollectAll() error = %v, wantErr %v", err, c.wantErr)
// 				return
// 			}
// 		})
// 	}

// }
