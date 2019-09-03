package db

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/garj4/lend/helpers"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// Defines test setup and teardown for the whole package
func TestMain(m *testing.M) {
	// Setup
	helpers.ConfigDir = "/tmp"

	exitCode := m.Run()

	// Teardown
	testDBPath := path.Join(helpers.ConfigDir, "sqlite.db")
	os.Remove(testDBPath)
	os.Exit(exitCode)
}

// A helper function to run a COUNT query to verify something exists
func exists(query string) (bool, error) {
	rows, err := database.query(query)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, err
	}

	count := 0
	err = rows.Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func TestAddPerson(t *testing.T) {
	t.Run("new person", func(_t *testing.T) {
		// Arrange
		testFirstName := "newPerson"

		// Act
		err := AddPerson(testFirstName, "lastName")

		// Assert
		assert.Nil(_t, err)

		personExists, err := exists(fmt.Sprintf("SELECT COUNT(1) FROM people WHERE firstname = \"%s\"", testFirstName))
		assert.Nil(_t, err)
		assert.True(_t, personExists)
	})

	t.Run("person already exists", func(_t *testing.T) {
		// Arrange
		testFirstName := "personAlreadyAdded"
		err := AddPerson(testFirstName, "lastName")
		assert.Nil(_t, err)

		// Act
		err = AddPerson(testFirstName, "lastName")

		// Assert
		assert.NotNil(_t, err)
	})
}

func TestAddTransaction(t *testing.T) {
	// Arrange
	testFirstName := "testTransactionPerson"
	testEvent := "testEvent"
	err := AddPerson(testFirstName, "lastName")
	assert.Nil(t, err)

	// Act
	err = AddTransaction(testEvent, testFirstName, 1.0)

	// Assert
	assert.Nil(t, err)

	transactionExists, err := exists(fmt.Sprintf("SELECT COUNT(1) FROM transactions WHERE event = \"%s\"", testEvent))
	assert.Nil(t, err)
	assert.True(t, transactionExists)
}

func TestPrintPeople(t *testing.T) {
	// Arrange
	names := []string{"test1", "printPeopleTest", "test3"}
	lastNames := []string{"lastName1", "otherLastName", "lastName3"}

	for i := 0; i < len(names); i++ {
		err := AddPerson(names[i], lastNames[i])
		assert.Nil(t, err)
	}

	testWriter := new(bytes.Buffer)

	// Act
	err := PrintPeople(testWriter)

	// Assert
	assert.Nil(t, err)

	for i := 0; i < len(names); i++ {
		expectedPerson := fmt.Sprintf(": %s %s\n", names[i], lastNames[i])
		assert.Contains(t, testWriter.String(), expectedPerson)
	}
}

func TestPrintTransactions(t *testing.T) {
	// Arrange
	testPerson := "printTransactionsPerson"
	err := AddPerson(testPerson, "lastName")
	assert.Nil(t, err)

	events := []string{"event1", "printTransactionTestEvent", "event3"}
	amounts := []float64{2.0, 1.34, -6.54}

	for i := 0; i < len(events); i++ {
		err := AddTransaction(events[i], testPerson, amounts[i])
		assert.Nil(t, err)
	}

	testWriter := new(bytes.Buffer)

	// Act
	err = PrintTransactions(testWriter)

	// Assert
	assert.Nil(t, err)

	for i := 0; i < len(events); i++ {
		expectedTransaction1 := fmt.Sprintf(": %s on date ", events[i])
		expectedTransaction2 := fmt.Sprintf(" for amount %.2f\n", amounts[i])
		assert.Contains(t, testWriter.String(), expectedTransaction1)
		assert.Contains(t, testWriter.String(), expectedTransaction2)
	}
}

func TestGetAmountOwed(t *testing.T) {
	t.Run("no transactions", func(_t *testing.T) {
		// Arrange
		testPerson := "nothingOwedPerson"
		err := AddPerson(testPerson, "lastName")
		assert.Nil(_t, err)

		// Act
		actualAmountOwed, err := GetAmountOwed(testPerson)

		// Assert
		assert.Nil(_t, err)
		assert.Equal(_t, 0.0, actualAmountOwed)
	})

	t.Run("owes money", func(_t *testing.T) {
		// Arrange
		testPerson := "owesMoneyPerson"
		err := AddPerson(testPerson, "lastName")
		assert.Nil(_t, err)

		amount1 := 6.53
		amount2 := -1.24

		err = AddTransaction("event", testPerson, amount1)
		assert.Nil(_t, err)

		err = AddTransaction("event", testPerson, amount2)
		assert.Nil(_t, err)

		// Act
		actualAmountOwed, err := GetAmountOwed(testPerson)

		// Assert
		assert.Nil(_t, err)
		assert.Equal(_t, amount1+amount2, actualAmountOwed)
	})
}
