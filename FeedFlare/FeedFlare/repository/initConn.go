package repository

import (
	"context"
	"feedflare/ent"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitConnection() *ent.Client {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err)
	}
	c := mysql.Config{
		DBName:    os.Getenv("MYSQL_DATABASE"),
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Addr:      os.Getenv("MYSQL_ADDR"),
		Net:       os.Getenv("NET_TYPE"),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	client, err := ent.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to mysql database: %v", err)
	}

	//defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
