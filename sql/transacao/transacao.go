package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO usuarios (nome) VALUES ($1, $2)")
	stmt.Exec(200, "Bia")
	stmt.Exec(201, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
