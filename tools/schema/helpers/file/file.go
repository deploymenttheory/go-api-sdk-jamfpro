package file

import (
	"fmt"
	"os"
)

// SaveStructsToFile saves the generated Go struct definitions to a specified file.
func SaveStructsToFile(structs string, filePath string) error {
	// Create or open the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write the struct definitions to the file
	_, err = file.WriteString(structs)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
