/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"novastar-cli/internal/client"
	conf "novastar-cli/internal/config"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func ExecuteInfo() {
	configManager := conf.NewConfigManager()
	AuthorizeCheck(configManager)

	config, err := configManager.ReadConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	url := fmt.Sprintf("https://%s:%d/terminal/core/v2/device/info", config.IpAddress, config.Port)
	method := "GET"

	data, err := client.Request(nil, url, method, config.Token)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	var infoResponse InfoResponse
	err = json.Unmarshal(data, &infoResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}
	fmt.Println(fmt.Sprint(infoResponse))
}

type InfoResponse struct {
	FPGAVersion     string `json:"fpga"`
	MainVersion     string `json:"main"`
	ProductName     string `json:"productName"`
	AtlasName       string `json:"atlasName"`
	SerialNumber    string `json:"sn"`
	RegisterAddress string `json:"registerAddress"`
	MacAddress      string `json:"mac"`
	PcbVersion      string `json:"pcbVersion"`
	AndroidVersion  string `json:"androidVersion"`
	OsVersion       string `json:"osVersion"`
}
