package dao

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var RecordNotFond = errors.New("RecordNotFond")

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "data/data.db")
	if err != nil {
		panic(err)
	}
}
