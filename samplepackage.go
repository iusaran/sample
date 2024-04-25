package sample

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
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

func UpdateFile(fileName string, newContent []byte) error {
    // Check if fileName is empty
    if fileName == "" {
        return errors.New("file name cannot be empty")
    }

    // Check if newContent is empty
    if len(newContent) == 0 {
        return errors.New("new content cannot be empty")
    }

    // Read the existing content of the file
    existingContent, err := ioutil.ReadFile(fileName)
    if err != nil {
        return err // Error reading the file
    }

    // Append the new content to the existing content
    updatedContent := append(existingContent, newContent...)

    // Write the updated content back to the file
    err = ioutil.WriteFile(fileName, updatedContent, 0644)
    if err != nil {
        return err // Error writing to the file
    }

    return nil // Update operation successful
}


func RenameFile(oldName, newName string) error {
    // Check if oldName or newName is empty
    if oldName == "" {
        return errors.New("old file name cannot be empty")
    }
    if newName == "" {
        return errors.New("new file name cannot be empty")
    }

    // Rename the file
    err := os.Rename(oldName, newName)
    if err != nil {
        return err // Error renaming the file
    }

    return nil // Rename operation successful
}

func DeleteFile(fileName string) error {
    // Check if fileName is empty
    if fileName == "" {
        return errors.New("file name cannot be empty")
    }

    // Attempt to remove the file
    err := os.Remove(fileName)
    if err != nil {
        return err // Error deleting the file
    }

    return nil // Deletion operation successful
}

// func MoveFile(oldPath, newPath string) error {
//     // Check if oldPath or newPath is empty
//     if oldPath == "" {
//         return errors.New("old path cannot be empty")
//     }
//     if newPath == "" {
//         return errors.New("new path cannot be empty")
//     }

//     // Move the file
//     err := os.Rename(oldPath, newPath)
//     if err != nil {
//         return err // Error moving the file
//     }

//     return nil // Move operation successful
// }

func MoveFile(oldPath, newPath string) error {
    // Check if the source file exists
    _, err := os.Stat(oldPath)
    if err != nil {
        if os.IsNotExist(err) {
            return fmt.Errorf("source file '%s' does not exist", oldPath)
        }
        return fmt.Errorf("error checking source file: %v", err)
    }

    // Check if the destination directory exists
    _, err = os.Stat(newPath)
    if err != nil {
        if os.IsNotExist(err) {
            // If the destination directory does not exist, create it
            err := os.MkdirAll(newPath, 0755)
            if err != nil {
                return fmt.Errorf("error creating destination directory: %v", err)
            }
        } else {
            return fmt.Errorf("error checking destination directory: %v", err)
        }
    }

    // Move the file
    err = os.Rename(oldPath, newPath)
    if err != nil {
        return fmt.Errorf("error moving file: %v", err)
    }

    return nil
}


func CopyFile(source, destination string) error {
    cmd := exec.Command("cmd", "/c", "Xcopy", source, destination)
    if err := cmd.Run(); err != nil {
        return err
    }
    return nil
}









