package main

import (
	"fmt"
	"novastar-cli/internal/fileexplorer"
)

// Example of how to use the file explorer
func main() {
	fmt.Println("=== File Explorer Demo ===")

	// Example 1: Quick selection from data folder
	fmt.Println("\n1. Quick file selection from data folder:")
	selectedFile, err := fileexplorer.SelectFileFromDataFolder()
	if err != nil {
		fmt.Printf("Error or cancelled: %v\n", err)
	} else {
		fmt.Printf("You selected: %s\n", selectedFile)
	}

	// Example 2: Custom directory with custom prompt
	fmt.Println("\n2. Custom directory selection with custom prompt:")
	customFile, err := fileexplorer.SelectFileWithPrompt("./data", "Please select a configuration file:")
	if err != nil {
		fmt.Printf("Error or cancelled: %v\n", err)
	} else {
		fmt.Printf("You selected: %s\n", customFile)
	}

	// Example 3: Using the FileExplorer directly for more control
	fmt.Println("\n3. Direct FileExplorer usage:")
	explorer := fileexplorer.NewFileExplorer()
	err = explorer.ScanDirectory("./data")
	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		return
	}

	fmt.Printf("Found %d files\n", explorer.GetFileCount())
	explorer.ListFiles()

	// You can also get specific files by index without prompting
	if explorer.GetFileCount() > 0 {
		firstFile, err := explorer.GetFileByIndex(1)
		if err == nil {
			fmt.Printf("First file: %s at %s\n", firstFile.Name, firstFile.Path)
		}
	}
}
