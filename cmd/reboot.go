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
	"novastar-cli/internal/response"
	"strings"

	"github.com/spf13/cobra"
)

// rebootCmd represents the reboot command
var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot the device",
	Long:  `Reboot the last novastar device you logged in to.`,
	Run: func(cmd *cobra.Command, args []string) {
		configManager := conf.NewConfigManager()
		ExecuteReboot(configManager)
	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)
}

func ExecuteReboot(configManager *conf.ConfigManager) {
	AuthorizeCheck(configManager)

	config, err := configManager.ReadConfig()
	url := fmt.Sprintf("https://%s:%d/terminal/core/v1/device/reboot", config.IpAddress, config.Port)
	method := "POST"

	payloadData := map[string]any{
		"reason": "manual reboot from CLI",
	}

	payloadDataBytes, err := json.Marshal(payloadData)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payloadDataBytes)))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", configManager.GetValue("token").(string))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	responseWrapper := response.ResponseWrapper{}
	json.Unmarshal(body, &responseWrapper)

	if responseWrapper.Code != 0 {
		fmt.Printf("Error rebooting device: %s\n", responseWrapper.Message)
		return
	}
	fmt.Println("Device rebooted successfully !")

}
