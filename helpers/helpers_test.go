package helpers

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	tmpPath := path.Join("/tmp", "testCreateFile")

	t.Run("File already exists", func(_t *testing.T) {
		// Arrange
		_, err := os.Create(tmpPath)

		assert.Nil(_t, err, "Failed to create tmp test file")

		// Act
		err = CreateFile(tmpPath)

		// Assert
		assert.Nil(_t, err)

		_, err = os.Stat(tmpPath)
		assert.Nil(_t, err)
	})

	t.Run("Creating new file", func(_t *testing.T) {
		// Arrange
		os.Remove(tmpPath)

		// Act
		err := CreateFile(tmpPath)

		// Assert
		assert.Nil(_t, err)

		_, err = os.Stat(tmpPath)
		assert.Nil(_t, err)
	})

	// Cleanup
	os.Remove(tmpPath)
}
