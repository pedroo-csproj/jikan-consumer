package env

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Host string
	Port string
)

func Load(filenames ...string) error {
	godotenv.Load()

	if value, found := os.LookupEnv("HOST"); found {
		Host = value

		fmt.Println(Host)
	} else {
		return errors.New("environment variable 'HOST' is required")
	}

	if value, found := os.LookupEnv("PORT"); found {
		Port = value

		fmt.Println(Host)
	} else {
		return errors.New("environment variable 'PORT' is required")
	}

	return nil
}
