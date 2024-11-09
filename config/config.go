package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	API APIConfig `mapstructure:"api"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Println(" config file:", err)

		return nil, err
	}

	return &config, nil
}
