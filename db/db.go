package db

import (
	"database/sql"
	"fmt"
	"path"

	"github.com/garj4/lend/helpers"
)

// Database represents the Database
type Database struct {
	init     bool
	dbDriver *sql.DB
}

var database Database

func (db *Database) initialize() error {
	if db.init {
		return nil
	}

	sqlitePath := path.Join(helpers.ConfigDir, "sqlite.db")

	err := helpers.CreateFile(sqlitePath)
	if err != nil {
		return err
	}

	db.dbDriver, err = sql.Open("sqlite3", sqlitePath)
	if err != nil {
		return err
	}

	// Create table for people
	statement, _ := db.dbDriver.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	// Create table for transactions
	// TODO: Reference a person from people table as an attribute
	statement, _ = db.dbDriver.Prepare("CREATE TABLE IF NOT EXISTS transactions (id INTEGER PRIMARY KEY, event TEXT, amount FLOAT, date DATETIME)")
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	db.init = true

	return nil
}

// AddRecord adds a new record into the database
func AddRecord(event, firstName string, amount float64) error {
	database.initialize()

	statement, _ := database.dbDriver.Prepare("INSERT INTO transactions (event, amount) VALUES (?, ?)")
	statement.Exec(event, fmt.Sprintf("%f", amount))

	return nil
}

// GetRecords returns the rows in the transactions table
func GetRecords() (*sql.Rows, error) {
	database.initialize()

	rows, err := database.dbDriver.Query("SELECT id, event, amount FROM transactions")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// TODO: Format package for JSON vs. terminal formatting
