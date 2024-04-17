package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DefaultConnectionString is the Default API connection String
	DefaultConnectionString = ""
	// DefaultApiPort is the Default API running Port
	DefaultApiPort = 0
	// DefaultSecretKey is the default secret key for Auth token assign
	DefaultSecretKey []byte
)

// Load loads all environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	DefaultApiPort, err = strconv.Atoi(os.Getenv("API_DEFAULT_PORT"))
	if err != nil {
		DefaultApiPort = 8080
	}

	DefaultConnectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DefaultSecretKey = []byte(os.Getenv("SECRET_KEY"))
}
