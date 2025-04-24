package main

import (
	"fmt"
	gedis "github.com/GedisCaching/Gedis/gedis"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadEnv loads environment variables with fallback to defaults
func LoadEnv() (addr string, password string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	addr = os.Getenv("ADDRESS")
	password = os.Getenv("PASSWORD")
	return addr, password
}

func main() {

	// Load environment variables
	Address, Password := LoadEnv()

	if Address == "" || Password == "" {
		log.Fatal("Missing required environment variables: ADDRESS and PASSWORD")
	}

	fmt.Printf("Connecting to Gedis at %s\n", Address)

	srv, err := gedis.NewGedis(gedis.Config{
		Address:  Address,
		Password: Password,
	})

	if err != nil {
		fmt.Println("Error creating server:", err)
		return
	}

	// Set a key-value pair
	srv.Set("name", "Gedis")

	// Get the value of the key
	value, exists := srv.Get("name")
	if !exists {
		fmt.Println("Key 'name' does not exist")
		return
	}

	// Print the value
	fmt.Println("Value of 'name':", value)
}
