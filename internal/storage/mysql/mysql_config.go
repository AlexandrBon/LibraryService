package mysql

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetDefaultConfig() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("couldn't recover information about config filename")
	}

	if err := godotenv.Load(filepath.Join(filename, "..", "..", "..", "..", "mysql.env")); err != nil {
		log.Fatal(err)
	}

	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    os.Getenv("MYSQL_NETWORK"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}

	fmt.Println(cfg.FormatDSN())
	return cfg.FormatDSN()
}
