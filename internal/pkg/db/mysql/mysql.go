package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDb() {
	db, err := sqlx.Open("mysql", "dharmasaputra:password@tcp(localhost)/gql_example")

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	DB = db
}
