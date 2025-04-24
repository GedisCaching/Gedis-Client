package main

import (
	"Gedis-Client/config"
	"Gedis-Client/operations"
	"fmt"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	fmt.Println("Initialization complete ✅")

	// Perform basic operations
	operations.BasicOperations()
	fmt.Println("Basic operations completed ✅")
}
