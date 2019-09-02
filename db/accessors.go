package db

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
)

// AddPerson adds a new person to the database if they have a unique first name
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

	_, err = database.dbDriver.Exec(addPersonQuery, firstName, lastName)
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

	formattedAmount := strconv.FormatFloat(amount, 'f', 2, 64)
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	formattedPersonID := strconv.Itoa(personID)

	_, err = database.dbDriver.Exec(addTransactionQuery, event, formattedAmount, formattedTime, formattedPersonID)
	if err != nil {
		return err
	}

	return nil
}

// PrintPeople prints the people table to the provided io.Writer
func PrintPeople(w io.Writer) error {
	err := database.initialize()
	if err != nil {
		return err
	}

	rows, err := database.dbDriver.Query(selectAllPeopleQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	var id int
	var firstName, lastName string
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			return fmt.Errorf("error when reading rows from DB: %s", err)
		}
		fmt.Fprintf(w, "%d: %s %s\n", id, firstName, lastName)
	}

	return nil
}

// PrintTransactions prints the transactions table to the provided io.Writer
func PrintTransactions(w io.Writer) error {
	err := database.initialize()
	if err != nil {
		return err
	}

	rows, err := database.dbDriver.Query(selectAllTransactionsQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	var id, person int
	var event, date string
	var amount float64
	for rows.Next() {
		err := rows.Scan(&id, &event, &amount, &date, &person)
		if err != nil {
			return fmt.Errorf("error when reading rows from DB: %s", err)
		}
		fmt.Fprintf(w, "%d: %s on date %s from personId %d for amount %.2f\n", id, event, date, person, amount)
	}

	return nil
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

	rows, err := database.dbDriver.Query(sumTransactionsQuery, personID)
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
