/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"novastar-cli/internal/file_explorer"
	"strings"

	"github.com/spf13/cobra"
)

// explorerCmd represents the explorer command
var explorerCmd = &cobra.Command{
	Use:   "explorer",
	Short: "Demonstrate file explorer functionality",
	Long: `A demonstration command that showcases the file explorer features.
This command allows you to:
- Browse files in the data directory
- Select files interactively
- See examples of different explorer usage patterns`,
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteExplorerDemo()
	},
}

func init() {
	rootCmd.AddCommand(explorerCmd)
}

func ExecuteExplorerDemo() {
	fmt.Println("🔍 File Explorer Demo")
	fmt.Println("====================")

	for {
		fmt.Println("\nChoose a demo option:")
		fmt.Println("[1] Quick file selection from data folder")
		fmt.Println("[2] Browse any directory")
		fmt.Println("[3] Advanced explorer features")
		fmt.Println("[4] Show file count only")
		fmt.Println("[q] Quit")
		fmt.Print("\nEnter your choice: ")

		var choice string
		fmt.Scanln(&choice)

		switch strings.ToLower(choice) {
		case "1":
			demoQuickSelection()
		case "2":
			demoBrowseDirectory()
		case "3":
			demoAdvancedFeatures()
		case "4":
			demoFileCount()
		case "q", "quit":
			fmt.Println("Goodbye! 👋")
			return
		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}
	}
}

func demoQuickSelection() {
	fmt.Println("\n📁 Quick File Selection Demo")
	fmt.Println("----------------------------")

	selectedFile, err := file_explorer.SelectFileFromDataFolder()
	if err != nil {
		fmt.Printf("❌ Selection failed or cancelled: %v\n", err)
	} else {
		fmt.Printf("✅ You selected: %s\n", selectedFile)
		fmt.Println("💡 This file path can now be used in your application logic!")
	}
}

func demoBrowseDirectory() {
	fmt.Println("\n📂 Browse Any Directory Demo")
	fmt.Println("----------------------------")

	fmt.Print("Enter directory path (or press Enter for './data'): ")
	var dirPath string
	fmt.Scanln(&dirPath)

	if dirPath == "" {
		dirPath = "./data"
	}

	selectedFile, err := file_explorer.SelectFileWithPrompt(dirPath, "Please select a file from this directory:")
	if err != nil {
		fmt.Printf("❌ Selection failed or cancelled: %v\n", err)
	} else {
		fmt.Printf("✅ You selected: %s\n", selectedFile)
	}
}

func demoAdvancedFeatures() {
	fmt.Println("\n⚙️  Advanced Explorer Features Demo")
	fmt.Println("----------------------------------")

	explorer := file_explorer.NewFileExplorer()

	fmt.Print("Enter directory to scan (or press Enter for './data'): ")
	var dirPath string
	fmt.Scanln(&dirPath)

	if dirPath == "" {
		dirPath = "./data"
	}

	fmt.Printf("🔍 Scanning directory: %s\n", dirPath)
	err := explorer.ScanDirectory(dirPath)
	if err != nil {
		fmt.Printf("❌ Error scanning directory: %v\n", err)
		return
	}

	fileCount := explorer.GetFileCount()
	fmt.Printf("📊 Found %d files\n\n", fileCount)

	if fileCount == 0 {
		fmt.Println("No files found in the directory.")
		return
	}

	// Show all files
	explorer.ListFiles()

	// Demonstrate getting file by index
	fmt.Println("\n🎯 Getting specific files by index:")
	for i := 1; i <= min(3, fileCount); i++ {
		file, err := explorer.GetFileByIndex(i)
		if err == nil {
			fmt.Printf("File #%d: %s (Path: %s)\n", file.Index, file.Name, file.Path)
		}
	}

	// Let user select a file
	fmt.Println("\n👆 Now you can select a file:")
	selectedFile, err := explorer.PromptForSelection()
	if err != nil {
		fmt.Printf("❌ Selection failed or cancelled: %v\n", err)
	} else {
		fmt.Printf("✅ Final selection: %s\n", selectedFile)
	}
}

func demoFileCount() {
	fmt.Println("\n📊 File Count Demo")
	fmt.Println("------------------")

	explorer := file_explorer.NewFileExplorer()

	fmt.Print("Enter directory to count files (or press Enter for './data'): ")
	var dirPath string
	fmt.Scanln(&dirPath)

	if dirPath == "" {
		dirPath = "./data"
	}

	err := explorer.ScanDirectory(dirPath)
	if err != nil {
		fmt.Printf("❌ Error scanning directory: %v\n", err)
		return
	}

	fileCount := explorer.GetFileCount()
	fmt.Printf("📁 Directory: %s\n", dirPath)
	fmt.Printf("📄 File count: %d\n", fileCount)

	if fileCount > 0 {
		fmt.Println("\n📋 File list (names only):")
		for i := 1; i <= fileCount; i++ {
			file, err := explorer.GetFileByIndex(i)
			if err == nil {
				fmt.Printf("  • %s\n", file.Name)
			}
		}
	}
}

// Helper function for min (since Go doesn't have a built-in min for int)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
