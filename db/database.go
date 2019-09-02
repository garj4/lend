package db

import (
	"database/sql"
	"errors"
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

	if err := db.openDB(); err != nil {
		return err
	}

	if err := db.createTables(); err != nil {
		return err
	}

	db.init = true

	return nil
}

func (db *Database) openDB() error {
	sqlitePath := path.Join(helpers.ConfigDir, "sqlite.db")

	err := helpers.CreateFile(sqlitePath)
	if err != nil {
		return err
	}

	db.dbDriver, err = sql.Open("sqlite3", sqlitePath)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) createTables() error {
	if _, err := db.dbDriver.Exec(createPeopleTableQuery); err != nil {
		return err
	}

	if _, err := db.dbDriver.Exec(createTransactionsTableQuery); err != nil {
		return err
	}

	return nil
}

func (db *Database) getPerson(firstName string) (int, error) {
	personID := -1

	rows, err := database.dbDriver.Query(selectPersonQuery, firstName)
	if err != nil {
		return personID, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&personID)
		return personID, err
	}

	// No one found
	return personID, errors.New("that person has not yet been added to the database. Use `lend addPerson <first name> <last name>` to add them")
}
