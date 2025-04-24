package config

import (
	"errors"
	gedis "github.com/GedisCaching/Gedis/gedis"
	"github.com/joho/godotenv"
	"os"
)

var (
	Address  string
	Password string

	// Gedis is the Gedis client instance
	Gedisclient *gedis.Gedis
)

// INIT initializes the configuration by loading environment variables
// and creating a Gedis client instance.
func Init() error {
	if err := godotenv.Load(".env"); err != nil {
		return errors.New("error loading .env file")
	}

	Address = os.Getenv("ADDRESS")
	Password = os.Getenv("PASSWORD")

	if Address == "" || Password == "" {
		return errors.New("missing required environment variables: ADDRESS and PASSWORD")
	}

	var err error
	Gedisclient, err = gedis.NewGedis(gedis.Config{
		Address:  Address,
		Password: Password,
	})

	if err != nil {
		return errors.New("failed to create Gedis client")
	}
	return nil
}
