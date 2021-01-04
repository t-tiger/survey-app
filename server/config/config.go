package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config is application config.
var Config appConfig

type appConfig struct {
	SecretKey        string   `mapstructure:"secret_key" validate:"required"`
	PostgresHost     string   `mapstructure:"postgres_host" validate:"required"`
	PostgresPort     int      `mapstructure:"postgres_port" validate:"numeric"`
	PostgresDB       string   `mapstructure:"postgres_db" validate:"required"`
	PostgresUser     string   `mapstructure:"postgres_user" validate:"required"`
	PostgresPassword string   `mapstructure:"postgres_password" validate:"required"`
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
}

func init() {
	v := viper.New()
	v.SetDefault("secret_key", "")
	v.SetDefault("postgres_host", "")
	v.SetDefault("postgres_port", 5432)
	v.SetDefault("postgres_db", "")
	v.SetDefault("postgres_user", "")
	v.SetDefault("postgres_password", "")
	v.SetDefault("allowed_origins", []string{})
	v.AutomaticEnv()
	if err := v.Unmarshal(&Config); err != nil {
		panic(err)
	}
}

// Validate validates the config values.
func (c *appConfig) Validate() error {
	return validator.New().Struct(c)
}
