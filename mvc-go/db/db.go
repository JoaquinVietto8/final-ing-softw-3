package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func StartDbEngine() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	var dsn string
	if dbHost == "/cloudsql/diario-428220:us-central1:diario" {
		dsn = fmt.Sprintf("%s:%s@unix(%s)/%s", dbUser, dbPass, dbHost, dbName)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	}

	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err = Db.Ping(); err != nil {
		panic(err)
	}
}
