package repository

import (
	"context"
	"database/sql"
	"insightstream/ent"
	"log"
	"log/slog"
	"os"
	"time"

	"entgo.io/ent/dialect"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var CoreDatabase *sql.DB

func InitConnection() *ent.Client {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err)
	}

	err = godotenv.Load(wd + "/.env")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err)
	}

	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbAddr := os.Getenv("MYSQL_ADDR")
	dbNet := os.Getenv("NET_TYPE")

	c := mysql.Config{
		DBName:               dbName,
		User:                 dbUser,
		Passwd:               dbPass,
		Addr:                 dbAddr,
		Net:                  dbNet,
		AllowNativePasswords: true,
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  utc,
	}

	CoreDatabase, err = sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to mysql database: %v", err)
	}

	clt, err := ent.Open(dialect.MySQL, c.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to mysql database: %v", err)
	}

	//defer client.Close()

	slog.Info("connected to mysql database")

	if err := clt.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	slog.Info("created schema resources")

	return clt
}
