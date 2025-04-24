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

	// Perform list operations
	operations.ListOperations()
	fmt.Println("List operations completed ✅")
	fmt.Println("")

	// Perform sorted set operations
	operations.SortedSetOperations()
	fmt.Println("Sorted set operations completed ✅")
	fmt.Println("")

	// Perform TTL operations
	operations.TTLOperations()
	fmt.Println("TTL operations completed ✅")
	fmt.Println("")

	// Perform hash operations
	operations.HashOperations()
	fmt.Println("Hash operations completed ✅")
	fmt.Println("")
}
