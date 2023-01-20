package repository

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"insightstream/ent"
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
		Loc:                  jst,
	}

	client, err := ent.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to mysql database: %v", err)
	}

	//defer client.Close()

	fmt.Println("connected to mysql database")

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("created schema resources")

	return client
}
