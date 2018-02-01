package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/backend/")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("bknd")

	viper.SetDefault("debug", false)
	viper.SetDefault("salt", "salt")
	viper.SetDefault("cookieDomain", "d.miaversa.com.br")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if viper.GetBool("debug") {
		log.SetFormatter(&log.TextFormatter{})
	}

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
}
