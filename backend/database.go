package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(ip_address:port)/db?parseTime=true")

	if err != nil {
		return nil, err
	}

	return db, nil
}