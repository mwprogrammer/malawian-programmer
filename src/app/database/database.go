package database

import (
	"database/sql"
	"fmt"
)

const (
	host = "192.186.200.247"
	port = 5432
	user = "go"
	password = "progress"
	dbname = "mwprogrammer"
)

var connection string

func Connect() error {

	connection = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connection)

	if err != nil { return err }

	defer db.Close()

	err = db.Ping()

	if err != nil { return err }

	fmt.Println("Connected!")

	return nil

}

