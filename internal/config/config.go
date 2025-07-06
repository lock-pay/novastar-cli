package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	IpAddress           string `json:"ip_address"`
	Port                int    `json:"port"`
	Token               string `json:"token"`
	ClientName          string `json:"client_name"`
	SerialNumber        string `json:"serial_number"`
	TokenExpirationDate string `json:"token_expiratio_date"` // Timestamp of the last login
}

type ConfigManager struct {
	viper *viper.Viper
}

func NewConfigManager() *ConfigManager {
	v := viper.New()

	v.SetDefault("username", "admin")
	v.SetDefault("password", "SN008@+")
	v.SetDefault("port", 16674)
	v.SetDefault("client_name", "NovaStar-CLI")

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	return &ConfigManager{
		viper: v,
	}
}

func (cm *ConfigManager) ReadConfig() (Config, error) {
	err := cm.viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = cm.viper.Unmarshal(&config)
	return config, err
}

func (cm *ConfigManager) WriteConfig() error {
	return cm.viper.WriteConfig()
}

func (cm *ConfigManager) SetValue(key string, value interface{}) {
	cm.viper.Set(key, value)
}

func (cm *ConfigManager) GetValue(key string) interface{} {
	return cm.viper.Get(key)
}

func (cm *ConfigManager) ShouldRelogin(config Config) bool {
	currentTime := time.Now()
	expirationDateStr := cm.viper.GetString("token_expiration_date")
	if expirationDateStr == "" {
		return true
	}
	expirationDate, err := time.Parse(time.RFC3339, expirationDateStr)
	if err != nil {
		return true
	}
	if currentTime.After(expirationDate) {
		return true
	}
	return false
}

func (cm *ConfigManager) CanLogin(config Config) bool {
	if config.Username == "" || config.Password == "" || config.IpAddress == "" || config.Port <= 0 {
		return false
	}
	if config.Token == "" || config.SerialNumber == "" {
		return false
	}
	return true
}

func (cm *ConfigManager) RefreshTokenExpiration() {
	cm.viper.Set("token_expiration_date", time.Now().Add(5*time.Minute).Format(time.RFC3339))
}
