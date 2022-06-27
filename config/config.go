package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration name
const (
	APIIP   = "API_IP"
	APIPORT = "API_PORT"
)

var defaultConfigurations = map[string]string{
	APIIP:   "localhost",
	APIPORT: "8888",
}

// Init initializes service configurations.
func Init(svcName string) {
	viper.SetEnvPrefix(svcName)
	loadConfig()
}

// Get gets environment variable.
func Get(name string) string {
	return viper.GetString(name)
}

func loadConfig() {
	loadDefaultConfig()
	loadFileConfig()
	loadEnvConfig()
}

func loadDefaultConfig() {
	for configKey, configValue := range defaultConfigurations {
		viper.SetDefault(configKey, configValue)
	}
}

func loadFileConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warn("Can't load config")
	}
}
func loadEnvConfig() {
	viper.AutomaticEnv()
}
