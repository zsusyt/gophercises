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
	//dbname = "zsusyt"
)

func main () {
	fmt.Println(dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	fmt.Println(dbname)
	err = createDB(db, dbname)
	//err = resetDB(db, dbname)
	must(err)
	must(createPhoneNumberTable(db))
	_, err = insertPhone(db, "1234567890")
	must(err)
	_, err = insertPhone(db, "123 456 7891")
	must(err)
	_, err = insertPhone(db, "(123) 456 7892")
	must(err)
	_, err = insertPhone(db, "(123) 456-7893")
	must(err)
	_, err = insertPhone(db, "123-456-7894")
	must(err)
	_, err = insertPhone(db, "123-456-7890")
	must(err)
	_, err = insertPhone(db, "1234567892")
	must(err)
	_, err = insertPhone(db, "(123)456-7892")
	must(err)
}

func insertPhone (db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createPhoneNumberTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)`
	_, err := db.Exec(statement)
	return err
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
