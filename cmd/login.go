/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"novastar-cli/internal/config"
	"novastar-cli/internal/response"
	"strings"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Accessing flags
		serial_Number, _ := cmd.Flags().GetString("sn")
		ipAddress, _ := cmd.Flags().GetString("ip")

		// Accessing arguments
		if len(args) > 0 {
			fmt.Printf("Arguments: %v\n", args)
		}

		// Pass flags to ExecuteLogin or handle them directly
		configManager := config.NewConfigManager()
		ExecuteLogin(configManager, serial_Number, ipAddress)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Adding flags
	loginCmd.Flags().String("ip", "", "IP address of the device")
	loginCmd.Flags().String("sn", "", "Port of the device")
}

func ExecuteLogin(configManager *config.ConfigManager, serial_Number, ipAddress string) {
	config, err := configManager.ReadConfig()
	url := fmt.Sprintf("https://%s:%d/terminal/core/v1/user/login", config.IpAddress, config.Port)
	method := "POST"

	payloadData := map[string]any{
		"sn":         serial_Number,
		"ip":         ipAddress,
		"username":   config.Username,
		"password":   config.Password,
		"loginType":  2,
		"clientId":   1,
		"clientName": config.ClientName,
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payloadBytes)))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

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

	responseWrapper := response.ResponseWrapper{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	if responseWrapper.Code != 0 {
		fmt.Printf("Login failed: %s\n", responseWrapper.Message)
		return
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(responseWrapper.Data, &loginResponse)
	if err != nil {
		fmt.Println("Error unmarshaling login response:", err)
		return
	}

	configManager.SetValue("token", loginResponse.Token)
	configManager.SetValue("serial_number", loginResponse.SerialNumber)
	configManager.SetValue("ip_address", ipAddress)
	configManager.SetValue("serial_number", loginResponse.SerialNumber)
	configManager.RefreshTokenExpiration()
}

type LoginResponse struct {
	Logined      bool   `json:"logined"`
	Password     string `json:"password"`
	SerialNumber string `json:"sn"`
	Token        string `json:"token"`
	Username     string `json:"username"`
	Validation   string `json:"validation"`
}
