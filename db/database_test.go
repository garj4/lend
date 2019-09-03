package db

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	// Act
	err := database.initialize()

	// Assert
	assert.Nil(t, err)
	assert.True(t, database.init)
}

func TestOpenDB(t *testing.T) {
	// Act
	err := database.openDB()
	assert.Nil(t, err)

	err = database.dbDriver.Ping()

	// Assert
	assert.Nil(t, err)
}

func TestCreateTables(t *testing.T) {
	// Arrange
	err := database.openDB()
	assert.Nil(t, err)

	// Act
	err = database.createTables()

	// Assert
	assert.Nil(t, err)

	_, err = database.query("SELECT * FROM people")
	assert.Nil(t, err)

	_, err = database.query("SELECT * FROM transactions")
	assert.Nil(t, err)
}

func TestGetPerson(t *testing.T) {
	t.Run("person is found", func(_t *testing.T) {
		// Arrange
		firstName := "testFirstName"
		lastName := "testLastName"

		_, err := database.exec("INSERT INTO people (firstname, lastname) VALUES (?, ?)", firstName, lastName)
		assert.Nil(_t, err)

		// Act
		actualID, err := database.getPerson(firstName)

		// Assert
		assert.Nil(_t, err)
		assert.True(_t, actualID > 0)
	})

	t.Run("person is not found", func(_t *testing.T) {
		// Act
		actualID, err := database.getPerson("notARealPerson")

		// Assert
		assert.NotNil(_t, err)
		assert.Equal(_t, -1, actualID)
	})
}
