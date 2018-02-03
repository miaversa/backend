package config

import (
	"github.com/spf13/viper"
)

// Load loads the configuration file using viper
func Load() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
