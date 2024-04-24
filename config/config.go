package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Dbconnect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1)/userdb")
	if err != nil {
		panic(err)
	}
	return db
}
