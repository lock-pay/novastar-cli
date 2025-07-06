/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"novastar-cli/internal/client"
	conf "novastar-cli/internal/config"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file to the device",
	Long:  `Upload a file to the last connected Novastar device.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	// Adding flags
	uploadCmd.Flags().StringP("file", "f", "", "Path to the file to upload")
	uploadCmd.Flags().StringP("target", "t", "", "Target directory on the device")

	uploadCmd.MarkFlagRequired("file")
	uploadCmd.MarkFlagRequired("target")
}

func ExecuteUpload(configManager *conf.ConfigManager, filePath string, targetPath string) {
	AuthorizeCheck(configManager)
	config, err := configManager.ReadConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	fileName := filepath.Base(filePath)
	fileExt := filepath.Ext(filePath)

	url := fmt.Sprintf("https://%s:%d/terminal/tools/v2/file/upload?fileName=%s&targetDir=%s", config.IpAddress, config.Port, fileName, targetPath)

	payload := strings.NewReader("<file contents here>")
	data, err := client.UploadFile(url, config.Token, payload, targetPath)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}
	fmt.Println("File uploaded successfully:", string(data))
	fmt.Printf("Uploaded file: %s (Extension: %s)\n", fileName, fileExt)
}
