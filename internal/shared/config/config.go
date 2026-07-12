package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Env  string
	}
	HTTP struct {
		Host string
		Port string
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
	Redis struct {
		Host string
		Port string
	}
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("APP_NAME", "cloud-native-ecommerce")
	v.SetDefault("APP_ENV", "development")
	v.SetDefault("HTTP_HOST", "0.0.0.0")
	v.SetDefault("HTTP_PORT", "8080")
	v.SetDefault("DATABASE_HOST", "localhost")
	v.SetDefault("DATABASE_PORT", "5432")
	v.SetDefault("DATABASE_USER", "postgres")
	v.SetDefault("DATABASE_PASS", "postgres")
	v.SetDefault("DATABASE_NAME", "ecommerce")
	v.SetDefault("REDIS_HOST", "localhost")
	v.SetDefault("REDIS_PORT", "6379")

	cfg := &Config{}
	cfg.App.Name = v.GetString("APP_NAME")
	cfg.App.Env = v.GetString("APP_ENV")
	cfg.HTTP.Host = v.GetString("HTTP_HOST")
	cfg.HTTP.Port = v.GetString("HTTP_PORT")
	cfg.Database.Host = v.GetString("DATABASE_HOST")
	cfg.Database.Port = v.GetString("DATABASE_PORT")
	cfg.Database.User = v.GetString("DATABASE_USER")
	cfg.Database.Pass = v.GetString("DATABASE_PASS")
	cfg.Database.Name = v.GetString("DATABASE_NAME")
	cfg.Redis.Host = v.GetString("REDIS_HOST")
	cfg.Redis.Port = v.GetString("REDIS_PORT")

	if cfg.HTTP.Port == "" {
		return nil, fmt.Errorf("HTTP_PORT is required")
	}

	return cfg, nil
}
