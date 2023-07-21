package configs

import (
	"os"
)

func GetEnv(key string) string {
	//var err = godotenv.Load(".env")
	/*if err != nil {
		panic("Error loading .env file")
	}*/
	return os.Getenv(key)
}
