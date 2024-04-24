package sample

import (
	"errors"
	"os"
)

// Add takes two integers and returns their sum.
func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}

func CreateFile(FileName string) (*os.File, error) {
	// Check if the file exists
	_, err := os.Stat(FileName)
	if err == nil {
		return nil, os.ErrExist // File already exists
	}
	if !os.IsNotExist(err) {
		return nil, err // Unexpected error occurred while checking file existence
	}

	// File doesn't exist, proceed to create it
	file, err := os.Create(FileName)
	if err != nil {
		return nil, err
	}

	// Ensure the file is closed when the function returns
	defer file.Close()

	return file, nil
}

func ReadFile(FileName string) ([]byte, error) {
	// Open the file
	file, err := os.Open(FileName)
	if err != nil {
		return nil, err // Error opening the file
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err // Error getting file info
	}
	fileSize := fileInfo.Size()

	// Read the file content into a byte slice
	content := make([]byte, fileSize)
	_, err = file.Read(content)
	if err != nil {
		return nil, err // Error reading the file
	}

	// Return the file content
	return content, nil
}

func WriteFile(fileName string, content []byte) error {
	// Check if the file name is empty
	if fileName == "" {
		return errors.New("file name cannot be empty")
	}

	// Check if the content is empty
	if len(content) == 0 {
		return errors.New("content cannot be empty")
	}

	// Open the file with write permissions, create if it doesn't exist, truncate if it does
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err // Error opening or creating the file
	}
	defer file.Close()

	// Write content to the file
	_, err = file.Write(content)
	if err != nil {
		return err // Error writing to the file
	}

	return nil // Write operation successful
}
