package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

var (
	appName string
	appPort int
)

// Load - Loads the application configurations
func Load() {
	viper.SetDefault("APP_NAME", "app")
	viper.SetDefault("APP_PORT", "8002")
	viper.SetDefault("GRPC_HOST", "")
	viper.SetDefault("GRPC_PORT", "")

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

// AppName - Application name
func AppName() string {
	if appName == "" {
		appName = ReadEnvString("APP_NAME")
	}
	return appName
}

// AppPort - Port that the application will run
func AppPort() int {
	if appPort == 0 {
		appPort = ReadEnvInt("APP_PORT")
	}
	return appPort
}

// ReadEnvString - Returns the value for the key from the configurations
func ReadEnvString(key string) string {
	checkIfSet(key)
	return viper.GetString(key)
}

// ReadEnvInt - Returns the value for the key from the configurations
func ReadEnvInt(key string) int {
	checkIfSet(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("Key %s is not a valid integer", key))
	}
	return v
}

func ReadEnvBool(key string) bool {
	checkIfSet(key)
	v, err := strconv.ParseBool(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("Key %s is not a valid boolean", key))
	}
	return v
}

func checkIfSet(key string) {
	if !viper.IsSet(key) {
		err := fmt.Errorf("Key %s is not set", key)
		panic(err)
	}
}
