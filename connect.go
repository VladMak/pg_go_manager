package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// ???????? ?????????? ? ??
	db, err := sql.Open("postgres", "user=postgres password=79911997 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ?????????? ??????? ? ????????? ???????????
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
}
