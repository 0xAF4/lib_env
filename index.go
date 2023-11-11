package lib_env

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
	}

	Http struct {
		Port  string `env:"HTTP_PORT"`
		Group string `env:"HTTP_GROUP"`
	}

	Jwt struct {
		Key       string `env:"JWT_KEY"`
		Algorithm string `env:"JWT_ALGORITHM"`
		LiveTime  string `env:"JWT_LIVE_TIME"`
	}

	Telegram struct {
		Token  string `env:"TELEGRAM_TOKEN"`
		ChatID string `env:"TELEGRAM_CHAT_ID"`
		Users  string `env:"TELEGRAM_USERS"`
	}

	Mail struct {
		Server     string `env:"MAIL_SERVER"`
		Port       string `env:"MAIL_PORT"`
		SenderMail string `env:"MAIL_SENDER_MAIL"`
	}
}

var (
	Con *Config
)

func init() {
	godotenv.Load()
	env.Parse(&Con)
}
