package adaptor

//import (
//	"errors"
//	"github.com/go-sql-driver/mysql"
//	"github.com/joho/godotenv"
//	"github.com/labstack/echo/v4"
//	"insightstream/ent"
//	"insightstream/ent/enttest"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"testing"
//	"time"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//func TestCollectAll(t *testing.T) {
//	jst, err := time.LoadLocation("Asia/Tokyo")
//	if err != nil {
//		os.Stderr.WriteString(err.Error())
//		panic(err)
//	}
//
//	wd, err := os.Getwd()
//	if err != nil {
//		os.Stderr.WriteString(err.Error())
//		panic(err)
//	}
//
//	err = godotenv.Load(wd + "/.env")
//	if err != nil {
//		os.Stderr.WriteString(err.Error())
//		panic(err)
//	}
//
//	dbName := os.Getenv("MYSQL_DATABASE")
//	dbUser := os.Getenv("MYSQL_USER")
//	dbPass := os.Getenv("MYSQL_PASSWORD")
//	dbAddr := os.Getenv("MYSQL_ADDR")
//	dbNet := os.Getenv("NET_TYPE")
//
//	cfg := mysql.Config{
//		DBName:               dbName,
//		User:                 dbUser,
//		Passwd:               dbPass,
//		Addr:                 dbAddr,
//		Net:                  dbNet,
//		AllowNativePasswords: true,
//		ParseTime:            true,
//		Collation:            "utf8mb4_unicode_ci",
//		Loc:                  jst,
//	}
//
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodGet, "/api/v2/fetch-feed/stored-all", nil)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//
//	c := e.NewContext(req, rec)
//
//	cases := map[string]struct {
//		input   echo.Context
//		wantErr bool
//	}{
//		"success": {
//			input:   c,
//			wantErr: false,
//		},
//		"failure": {
//			input:   nil,
//			wantErr: true,
//		},
//	}
//
//	cl := enttest.Open(t, "mysql", cfg.FormatDSN())
//	defer func(cl *ent.Client) {
//		err := cl.Close()
//		if err != nil {
//			t.Errorf("error: %v", err)
//		}
//	}(cl)
//
//	for name, c := range cases {
//		t.Run(name, func(t *testing.T) {
//			err := CollectAll(c.input, cl)
//
//			if errors.Is(err, errors.New("context is nil")) {
//				t.Errorf("CollectAll() got = %v, want %v", err, errors.New("context is nil"))
//				return
//
//			}
//
//			if err != nil {
//				t.Errorf("error: %v", err)
//			}
//		})
//	}
//}
