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

	stmt, _ := db.Prepare("UPDATE usuarios SET nome = $1 WHERE id = $2")

	stmt.Exec("Uóxiton Istive", 3)
	stmt.Exec("Sheristone Ualeska", 4)

	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = $1")

	stmt2.Exec(12)
}
