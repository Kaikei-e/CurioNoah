package repository

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"

	"entdemo/ent"

	_ "github.com/go-sql-driver/mysql"
)

func InitConnection() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		os.Stderr.WriteString(err.Error())

	}
	c := mysql.Config{
		DBName:    "db",
		User:      "user",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	client, err := ent.Open()
	if err != nil {
		log.Fatalf("failed opening connection to mysql database: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
