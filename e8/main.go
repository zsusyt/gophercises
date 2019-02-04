package main

import (
	"bytes"
	"database/sql"
	"fmt"

	_"github.com/lib/pq"
)


const (
	host = "localhost"
	port = 5432
	user = "zsusyt"
	password = ""
	dbname = "gophercises_phone"
)

func main () {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)

	err = resetDB(db, dbname)
	must(err)
	db.Close()

	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func must (err error) {
	if err != nil {
		panic(err)
	}
}

func resetDB(db *sql.DB, name string) error{
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
