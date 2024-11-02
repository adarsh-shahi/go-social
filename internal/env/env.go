package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetString(key, fallback string) string {
	log.Println(key)
	log.Println(os.Getenv(key))
	val, ok := os.LookupEnv(key)
	log.Println(val, ok)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)

	if err != nil {
		return fallback
	}
	return valAsInt
}
