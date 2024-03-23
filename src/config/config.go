package config

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	// DefaultConnectionString is the Default API connection String
	DefaultConnectionString = ""
	// DefaultApiPort is the Default API running Port
	DefaultApiPort = 0
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
