package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	con := "user=storeUser dbname=postgres password=localdb host=localhost sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err.Error())
	}
	return db
}