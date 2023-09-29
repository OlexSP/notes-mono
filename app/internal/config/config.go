package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
	"sync"
	"time"
)

// SQLDateFormat - date format for sql struct fields
const SQLDateFormat = time.RFC3339

// Config - config
type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"local"`
	HTTP     struct {
		IP           string        `yaml:"ip" env:"HTTP-IP"`
		Port         int           `yaml:"port" env:"HTTP-PORT"`
		ReadTimeout  time.Duration `yaml:"read-timeout" env:"HTTP-READ-TIMEOUT"`
		WriteTimeout time.Duration `yaml:"write-timeout" env:"HTTP-WRITE-TIMEOUT"`
		CORS         struct {
			AllowedMethods     []string `yaml:"allowed_methods" env:"HTTP-CORS-ALLOWED-METHODS"`
			AllowedOrigins     []string `yaml:"allowed_origins"`
			AllowCredentials   bool     `yaml:"allow_credentials"`
			AllowedHeaders     []string `yaml:"allowed_headers"`
			OptionsPassthrough bool     `yaml:"options_passthrough"`
			ExposedHeaders     []string `yaml:"exposed_headers"`
			Debug              bool     `yaml:"debug"`
		} `yaml:"cors"`
	} `yaml:"http"`
	AppConfig struct {
		LogLevel  string `yaml:"log-level" env:"LOG_LEVEL" env-default:"trace"`
		AdminUser struct {
			Email    string `yaml:"email" env:"ADMIN_EMAIL" env-default:"admin"`
			Password string `yaml:"password" env:"ADMIN_PWD" env-default:"admin"`
		} `yaml:"admin"`
	} `yaml:"app"`
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
