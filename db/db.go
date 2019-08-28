package db

import (
	"database/sql"
	"errors"
	"fmt"
	"path"
	"strconv"
	"time"

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
	statement, _ = db.dbDriver.Prepare("CREATE TABLE IF NOT EXISTS transactions (id INTEGER PRIMARY KEY, event TEXT, amount FLOAT, date DATETIME, person INTEGER)")
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	db.init = true

	return nil
}

func (db *Database) getPerson(firstName string) (int, error) {
	personID := -1

	query := fmt.Sprintf("SELECT id FROM people WHERE firstname = \"%s\"", firstName)

	rows, err := database.dbDriver.Query(query)
	if err != nil {
		return personID, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&personID)
		return personID, err
	}

	// No one found
	return personID, errors.New("that person has not yet been added to the database. Use `lend add-person <first name> <last name>` to add them")
}

// AddPerson adds a new person to the database, if they don't already exist
func AddPerson(firstName, lastName string) error {
	err := database.initialize()
	if err != nil {
		return err
	}

	// Check if the person already exists
	personID, err := database.getPerson(firstName)
	if personID > 0 || err == nil {
		return errors.New("someone with that firstname has already been added to the database")
	}

	statement, err := database.dbDriver.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(firstName, lastName)
	if err != nil {
		return err
	}

	return nil
}

// AddTransaction adds a new transaction into the database
func AddTransaction(event, firstName string, amount float64) error {
	err := database.initialize()
	if err != nil {
		return err
	}

	personID, err := database.getPerson(firstName)
	if err != nil {
		return err
	}

	statement, err := database.dbDriver.Prepare("INSERT INTO transactions (event, amount, date, person) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(event, fmt.Sprintf("%f", amount), time.Now().Format("2006-01-02 15:04:05"), strconv.Itoa(personID))
	if err != nil {
		return err
	}

	return nil
}

// GetRecords returns the rows in the transactions table
func GetRecords() (*sql.Rows, error) {
	err := database.initialize()
	if err != nil {
		return nil, err
	}

	rows, err := database.dbDriver.Query("SELECT id, event, amount, date, person FROM transactions")
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// GetAmountOwed returns the amount a single person owes
func GetAmountOwed(name string) (float64, error) {
	err := database.initialize()
	if err != nil {
		return 0, err
	}

	personID, err := database.getPerson(name)
	if err != nil {
		return 0, err
	}

	rows, err := database.dbDriver.Query(fmt.Sprintf("SELECT SUM(amount) FROM transactions WHERE person = %d", personID))
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	sum := 0.0
	if rows.Next() {
		err := rows.Scan(&sum)
		return sum, err
	}

	return 0, errors.New("error summing amount owed")
}

// TODO: Format package for JSON vs. terminal formatting
