package appconfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error happened while loading environment...%v", err)
	}
	return &Env{
		DB_HOST:             os.Getenv("DB_HOST"),
		DB_USER:             os.Getenv("DB_USER"),
		DB_PASSWORD:         os.Getenv("DB_PASSWORD"),
		DB_NAME:             os.Getenv("DB_NAME"),
		PORT:                os.Getenv("PORT"),
		SECRET_KEY:          os.Getenv("SECRET_KEY"),
		SECRET_ACCESSKEY:    os.Getenv("SECRET_ACCESSKEY"),
		SECRET_REFRESHKEY:   os.Getenv("SECRET_REFRESHKEY"),
		EXPIRATION_TIME_SEC: os.Getenv("EXPIRATION_TIME_SEC"),
	}, nil
}
