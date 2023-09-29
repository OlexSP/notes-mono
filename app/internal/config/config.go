package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
	"sync"
)

type Config struct {
	LogLevel      string `env:"LOG_LEVEL" envDefault:"local"`
	IsDebug       bool   `env:"IS_DEBUG" envDefault:"false"`
	IsDevelopment bool   `env:"IS_DEV" envDefault:"false"`
	Listen        struct {
		Type   string `env:"LISTEN_TYPE" envDefault:"port"`
		BindIP string `env:"BIND_IP" envDefault:"0.0.0.0"`
		Port   string `env:"PORT" envDefault:"10000"`
	}
	AppConfig struct {
		LogLevel  string //`env:"LOG_LEVEL" envDefault:"info"`
		AdminUser struct {
			Email    string `env:"ADMIN_USER_EMAIL" env-required:"true"`
			Password string `env:"ADMIN_PWD" env-required:"true"`
		}
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		slog.Info("GetConfig")
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "notes-mono"
			description, _ := cleanenv.GetDescription(instance, &helpText)
			slog.Info("cannot read env", slog.Any("error:", err), slog.Any("error:", description))
			os.Exit(1)
		}
	})
	return instance
}
