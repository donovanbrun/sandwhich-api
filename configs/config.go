package configs

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(key string) string {
	var err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}
