package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariable struct {
	PORT          string
	JWT_SECRET    string
	DATABASE_URL  string
	DATABASE_NAME string
}

var Env = &EnvVariable{}

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	Env = &EnvVariable{
		PORT:          os.Getenv("PORT"),
		JWT_SECRET:    os.Getenv("JWT_SECRET"),
		DATABASE_URL:  os.Getenv("DATABASE_URL"),
		DATABASE_NAME: os.Getenv("DATABASE_NAME"),
	}
}
