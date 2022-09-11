package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

func NewDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:password@tcp(127.0.0.1:3306)/link_management?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
