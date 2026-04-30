// Package configs provides configuration for the application.
package configs

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds the application configuration.
type Config struct {
	Port  int    `envconfig:"PORT"  default:"4000"`
	Stage string `envconfig:"STAGE" default:"dev"  validate:"oneof=dev stg prod"`

	// PGUser     string `envconfig:"PG_USER"     required:"true"`
	// PGPassword string `envconfig:"PG_PASSWORD" required:"true"`
	// PGHost     string `envconfig:"PG_HOST"     required:"true"`
	// PGPort     int    `envconfig:"PG_PORT"     default:"5432"`
	// PGDB       string `envconfig:"PG_DB"       required:"true"`

	// JWTSecret string `envconfig:"JWT_SECRET" required:"true" validate:"min=32"`
	// JWTExpiry string `envconfig:"JWT_EXPIRY" default:"4h"`

	// RedisHost     string `envconfig:"REDIS_HOST"     default:"localhost"`
	// RedisPort     int    `envconfig:"REDIS_PORT"     default:"6379"`
	// RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`
	// RedisDB       int    `envconfig:"REDIS_DB"       default:"0"`
}

// func (c *Config) PGUrl() string {
//     return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
//         url.QueryEscape(c.PGUser),
//         url.QueryEscape(c.PGPassword),
//         c.PGHost, c.PGPort, c.PGDB,
//     )
// }

// func (c *Config) RedisUrl() string {
//     return fmt.Sprintf("redis://%s:%d/%d", c.RedisHost, c.RedisPort, c.RedisDB)
// }

// Load loads the configuration from environment variables.
func Load() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}

	return &cfg, nil
}
