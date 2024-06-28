package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	ENV        string
	DB_URL     string
	JWT_SECRET string
}

var Get *Config

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := &Config{
		ENV:        viper.GetString("ENV"),
		DB_URL:     viper.GetString("DB_URL"),
		JWT_SECRET: viper.GetString("JWT_SECRET"),
	}
	Get = config
	return config, nil
}

// IsProd checks if the environment is set to "production".
//
// This function takes a pointer to a Config struct as its receiver.
// It returns a boolean value indicating whether the environment is set to "production".
func (c *Config) IsProd() bool {
	return c.ENV == "production"
}

// IsDev checks if the configuration environment is set to "development".
//
// No parameters.
// Returns a boolean value.
func (c *Config) IsDev() bool {
	return c.ENV == "development"
}

func (c *Config) IsValid() error {
	if c.ENV != "development" && c.ENV != "production" {
		return errors.New("invalid env")
	}

	if c.DB_URL == "" {
		return errors.New("missing db url")
	}

	if c.JWT_SECRET == "" {
		return errors.New("missing jwt secret")
	}
	return nil
}
