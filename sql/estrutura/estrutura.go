package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cursogo sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
