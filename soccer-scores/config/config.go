package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App    App    `yaml:"app"`
	Secret string `yaml:"secret"`
}

type App struct {
	Name string `yaml:"name"`
}

func (c *Config) Parse() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}
