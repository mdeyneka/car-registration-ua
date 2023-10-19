package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DbHost     string `env:"DB_HOST,required"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbName     string `env:"DB_NAME,required"`
	DbPort     int    `env:"DB_PORT,required"`
}

var Dbconfig Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env file: %e", err)
	}

	err = env.Parse(&Dbconfig)
	if err != nil {
		log.Fatalf("Unable to parse ennvironment variables: %e", err)
	}
}
