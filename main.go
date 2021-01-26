package main

import (
	"fmt"
	"log"
	_ "github.com/lib/pq"

	"database/sql"
)

func main() {
	//создаем подключение к базе данных и сохраняем в db
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost/gobase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type note struct {
		id   int
		text string
	}

	notes := []note{}

	for rows.Next() {
		p := note{}
		err := rows.Scan(&p.id, &p.text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		notes = append(notes, p)

		for _, p := range notes {
			fmt.Println(p.id, p.text)
		}
	}

}

type note struct {
	id   int
	text string
}
