package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	extInt, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	return &Config{
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
			Host: os.Getenv("SERVER_HOST"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Name: os.Getenv("DB_NAME"),
			Pass: os.Getenv("DB_PASS"),
			Tz:   os.Getenv("DB_TZ"),
		},
		Jwt: Jwt{
			Key: os.Getenv("JWT_KEY"),
			Exp: extInt,
		},
	}
}
