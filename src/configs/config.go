package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvVar struct {
	PORT      string
	SECRET    string
	MONGO_URL string
	DATABASE  string
}

var Env = EnvVar{
	PORT:      GetEnv("PORT"),
	SECRET:    GetEnv("SECRET"),
	MONGO_URL: GetEnv("MONGO_URL"),
	DATABASE:  GetEnv("DATABASE"),
}

func GetEnv(key string) string {
	var err = godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file, reading system environment variables")
	}
	return os.Getenv(key)
}
