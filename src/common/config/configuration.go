package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// New create new configuration object
func New() (*Configuration, error) {
	viper.SetConfigName(".env")
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	cfg := new(Configuration)
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}
