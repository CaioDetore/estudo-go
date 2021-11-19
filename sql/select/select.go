package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cursogo ")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, _ := db.Query("SELECT id, nome FROM usuarios WHERE id < $1", 5)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

}
