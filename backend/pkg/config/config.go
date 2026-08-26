package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("警告: .envファイルが見つかりませんでした")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
