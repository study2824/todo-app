package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
	"log"
	"os"
)

var DB *sql.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: loading .env file")
	}

	tmp, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	DB = tmp
}
