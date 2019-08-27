package helpers

import "os"

// CreateFile creates a file if it doesn't already exist
func CreateFile(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(path)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}
