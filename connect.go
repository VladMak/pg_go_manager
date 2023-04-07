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
	rows, err := db.Query("SELECT * FROM pg_database")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var datname, datlocprovider, datistemplate, datallowconn, datcollate, datctype string
		var datdba, encoding, datconnlimit, datfrozenxid, datminmxid, dattablespace int
		var daticulocale, atcollversion, datacl *string

		err = rows.Scan(&id, &datname, &datdba, &encoding, &datlocprovider, &datistemplate, &datallowconn, &datconnlimit, &datfrozenxid, &datminmxid, &dattablespace, &datcollate, &datctype, &daticulocale, &atcollversion, &datacl)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id=%d, datname=%s, datdba=%d, encoding=%d, datlocprovider=%s, datistemplate=%s, datallowconn=%s, datconnlimit=%d, datfrozenxid=%d, datminmxid=%d, dattablespace=%d, datcollate=%s, datctype=%s, daticulocale=%v, datcollversion=%v, atcollversion=%s, datacl=%v\n", id, datname, datdba, encoding, datlocprovider, datistemplate, datallowconn, datconnlimit, datfrozenxid, datminmxid, dattablespace, datcollate, datctype, daticulocale, atcollversion, datacl)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
}
