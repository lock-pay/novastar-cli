package fileexplorer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FileInfo represents a file with its index and path
type FileInfo struct {
	Index int
	Name  string
	Path  string
}

// FileExplorer handles file exploration functionality
type FileExplorer struct {
	files []FileInfo
}

// NewFileExplorer creates a new FileExplorer instance
func NewFileExplorer() *FileExplorer {
	return &FileExplorer{
		files: make([]FileInfo, 0),
	}
}

// ScanDirectory scans the specified directory for files and returns them with indexes
func (fe *FileExplorer) ScanDirectory(dirPath string) error {
	fe.files = make([]FileInfo, 0)

	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dirPath)
	}

	index := 1
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, only include files
		if !info.IsDir() {
			relPath, _ := filepath.Rel(dirPath, path)
			fe.files = append(fe.files, FileInfo{
				Index: index,
				Name:  relPath,
				Path:  path,
			})
			index++
		}
		return nil
	})

	return err
}

// ListFiles displays all found files with their indexes
func (fe *FileExplorer) ListFiles() {
	if len(fe.files) == 0 {
		fmt.Println("No files found in the specified directory.")
		return
	}

	fmt.Println("Files found:")
	fmt.Println("=============")
	for _, file := range fe.files {
		fmt.Printf("[%d] %s\n", file.Index, file.Name)
	}
	fmt.Println("=============")
}

// PromptForSelection prompts the user to select a file by index
func (fe *FileExplorer) PromptForSelection() (string, error) {
	if len(fe.files) == 0 {
		return "", fmt.Errorf("no files available for selection")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter the number of the file you want to select (1-%d) or 'q' to quit: ", len(fe.files))

		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("error reading input: %v", err)
		}

		input = strings.TrimSpace(input)

		// Allow user to quit
		if strings.ToLower(input) == "q" || strings.ToLower(input) == "quit" {
			return "", fmt.Errorf("user cancelled selection")
		}

		// Try to parse the number
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		// Validate the choice
		if choice < 1 || choice > len(fe.files) {
			fmt.Printf("Invalid choice. Please enter a number between 1 and %d.\n", len(fe.files))
			continue
		}

		// Return the selected file path
		selectedFile := fe.files[choice-1]
		fmt.Printf("Selected: %s\n", selectedFile.Name)
		return selectedFile.Path, nil
	}
}

// ExploreAndSelect is a convenience method that combines scanning, listing, and selection
func (fe *FileExplorer) ExploreAndSelect(dirPath string) (string, error) {
	// Scan the directory
	err := fe.ScanDirectory(dirPath)
	if err != nil {
		return "", err
	}

	// List the files
	fe.ListFiles()

	// Prompt for selection
	return fe.PromptForSelection()
}

// GetFileCount returns the number of files found
func (fe *FileExplorer) GetFileCount() int {
	return len(fe.files)
}

// GetFileByIndex returns the file info for a specific index (1-based)
func (fe *FileExplorer) GetFileByIndex(index int) (FileInfo, error) {
	if index < 1 || index > len(fe.files) {
		return FileInfo{}, fmt.Errorf("invalid index: %d", index)
	}
	return fe.files[index-1], nil
}

// QuickFileSelection is a utility function for quick file selection from a directory
// This is the main function other commands should use
func QuickFileSelection(dirPath string) (string, error) {
	explorer := NewFileExplorer()
	return explorer.ExploreAndSelect(dirPath)
}

// SelectFileFromDataFolder is a convenience function specifically for the data folder
func SelectFileFromDataFolder() (string, error) {
	dataPath := "./data"
	return QuickFileSelection(dataPath)
}

// SelectFileWithPrompt allows custom prompt message
func SelectFileWithPrompt(dirPath, promptMessage string) (string, error) {
	explorer := NewFileExplorer()

	err := explorer.ScanDirectory(dirPath)
	if err != nil {
		return "", err
	}

	explorer.ListFiles()

	if promptMessage != "" {
		fmt.Println(promptMessage)
	}

	return explorer.PromptForSelection()
}
