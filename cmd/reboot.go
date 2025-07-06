/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"novastar-cli/internal/client"
	conf "novastar-cli/internal/config"

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

	_, err = client.Request(payloadData, url, method, config.Token)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	fmt.Println("Device rebooted successfully !")

}
