package env

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	Host     string
	Port     string
	JikanApi string
)

func Load(filenames ...string) error {
	godotenv.Load()

	if value, found := os.LookupEnv("HOST"); found {
		Host = value
	} else {
		return errors.New("environment variable 'HOST' is required")
	}

	if value, found := os.LookupEnv("PORT"); found {
		Port = value
	} else {
		return errors.New("environment variable 'PORT' is required")
	}

	if value, found := os.LookupEnv("JIKAN_API"); found {
		JikanApi = value
	} else {
		return errors.New("environment variable 'JIKAN_API' is required")
	}

	return nil
}
