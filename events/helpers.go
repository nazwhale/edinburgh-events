package events

import (
	"fmt"
	"os"
)

// writeToCSV takes CSV content as a string and writes it to a specified file.
// It returns an error if the file cannot be created or written to.
func writeToCSV(csvContent, fileName string) error {
	// Create or overwrite the file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Write the CSV content to the file
	_, err = file.WriteString(csvContent)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil // No error occurred
}
