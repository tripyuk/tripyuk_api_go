package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

//var (
//	config       *Config
//	cfgLock      sync.Once
//	cfgType      string = "yml"
//	cfgPath      string = "."
//	cfgName      string = ".env"
//	cfgEnvPrefix string = "AGT"
//)

// New create new configuration object
func New() (*Configuration, error) {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")
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
