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
	fmt.Println("")

	// Perform basic operations
	operations.BasicOperations()
	fmt.Println("Basic operations completed ✅")
	fmt.Println("")


	// Perform numeric operations
	operations.NumericOperations()
	fmt.Println("Numeric operations completed ✅")
	fmt.Println("")
}
