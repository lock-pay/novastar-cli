/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", config.Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var infoResponse InfoResponse
	err = json.Unmarshal(body, &infoResponse)
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
