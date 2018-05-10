package dao

import (
	"database/sql"

	"sono/log"

	_ "github.com/go-sql-driver/mysql"
)

// Dao .
type Dao struct {
	db *sql.DB
}

// New .
func New() (d *Dao) {
	var err error
	d = &Dao{}
	d.db, err = sql.Open("mysql", "127.0.0.1:3306")

	if err != nil {
		log.Error("sql.Open()", err)
	}
	return
}
