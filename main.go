package main

import (
	"Gedis-Client/config"
	"fmt"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	fmt.Println("Initialization complete.")

	// Example usage of the Gedis client
	config.Gedisclient.Set("key", "value")

	value, exists := config.Gedisclient.Get("key")

	if !exists {
		log.Fatalf("Key does not exist")
	}

	fmt.Println("Retrieved value:", value)
}
