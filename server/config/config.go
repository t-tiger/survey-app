package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config is application config.
var Config appConfig

type appConfig struct {
	PostgresHost     string `mapstructure:"postgres_host" validate:"required"`
	PostgresPort     int    `mapstructure:"postgres_port" validate:"numeric"`
	PostgresDB       string `mapstructure:"postgres_db" validate:"required"`
	PostgresUser     string `mapstructure:"postgres_user" validate:"required"`
	PostgresPassword string `mapstructure:"postgres_password" validate:"required"`
}

func init() {
	v := viper.New()
	v.SetDefault("postgres_host", "")
	v.SetDefault("postgres_port", 5432)
	v.SetDefault("postgres_db", "")
	v.SetDefault("postgres_user", "")
	v.SetDefault("postgres_password", "")
	v.AutomaticEnv()
	if err := v.Unmarshal(&Config); err != nil {
		panic(err)
	}
}

// Validate validates the config values.
func (c *appConfig) Validate() error {
	return validator.New().Struct(c)
}
