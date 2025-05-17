package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

const (
	host = "192.186.200.247"
	port = 5432
	user = "go"
	password = "progress"
	dbname = "mwprogrammer"
)

var connection string
var db *sql.DB

func Connect() error {

	connection = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", connection)

	if err != nil { return err }

	err = db.Ping()

	if err != nil { return err }

	fmt.Println("Connected!")

	return nil

}

func Disconnect() error {

	err := db.Close()

	if err != nil { return err }

	fmt.Println("Disconnected!")

	return nil

}

func GetRecords(query string) (*sql.Rows, error) {
	
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	return rows, nil

}

// Adds a new record to a table in a database.
func AddRecord(record any) error {

	entity := reflect.TypeOf(record)
	properties := reflect.ValueOf(record)

	table := entity.Name() + "s"
	fields := make([]string, properties.NumField())
	values := make([]any, properties.NumField())

	for i := range fields {
		fields = append(fields, entity.Field(i).Name)
	}

	for i := range values {
		values = append(values, properties.FieldByName(fields[i]))
	}

	names := strings.TrimSuffix(strings.Join(fields, ", "), ",")
	count := strings.TrimSuffix(strings.Repeat("?, ", len(fields)), ", ")
	
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, names, count)

	_, err := db.Exec(query, values...)

	if err != nil {
		return err
	}

	return nil

}

func DeleteRecord(record any) (sql.Result, error) {
	
	entity := reflect.TypeOf(record)
	properties := reflect.ValueOf(record)

	table := entity.Name() + "s"
	fields := make([]string, properties.NumField())
	values := make([]any, properties.NumField())

	for i := range fields {
		fields = append(fields, entity.Field(i).Name)
	}

	for i := range values {
		values = append(values, properties.FieldByName(fields[i]))
	}

	names := strings.TrimSuffix(strings.Join(fields, ", "), ",")
	count := strings.TrimSuffix(strings.Repeat("?, ", len(fields)), ", ")
	
	query := fmt.Sprintf("DELETE FROM %s (%s) VALUES (%s)", table, names, count)

	results, err := db.Exec(query, values...)

	if err != nil {
		return nil, err
	}

	return results, nil

}
