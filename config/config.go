package config

import (
	"github.com/spf13/viper"
	cli "gopkg.in/urfave/cli.v1"
)

type Config struct {
	Database Database
}

type Database struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

// New configuration from file or from cli / env
func New(c *cli.Context) (*Config, error) {
	cfg, err := LoadConfigFile(c.String("config"))
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// LoadConfigFile loads any configuration file from JSON, YAML, TOML
func LoadConfigFile(path string) (*Config, error) {
	config := Config{}

	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
