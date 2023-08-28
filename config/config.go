package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	SessionEXP int
	JWTSecret  string
}

var cfg *Config

func SetConfig() *Config {
	var err error
	cfg = &Config{}

	cfg.SessionEXP, err = strconv.Atoi(os.Getenv("SESSION_EXP"))
	if err != nil {
		log.Fatalln("Session exp not valid")
	}

	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	if cfg.JWTSecret == "" {
		log.Fatalln("Jwt secret not valid")
	}

	return cfg
}

func GetConfig() *Config {
	return cfg
}
