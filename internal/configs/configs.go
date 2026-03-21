package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kaluginivann/Aegis/internal/logger"
)

type Config struct {
	FilePath string
	Logger   logger.Interface
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default settiongs")
	}
	return &Config{
		FilePath: os.Getenv("FILE_PATH"),
		Logger:   logger.NewLogger(),
	}
}
