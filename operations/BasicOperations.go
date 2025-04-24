package operations

import (
	"Gedis-Client/config"
	"fmt"
)

// BasicOperations demonstrates basic operations using the Gedis client.
func BasicOperations() {
	// Example usage of the Gedis client
	config.GedisClient.Set("key", "value")

	// Retrieve the value
	value, _ := config.GedisClient.Get("key")
	fmt.Println("Retrieved value:", value)

	// Delete the key
	config.GedisClient.Delete("key")
	fmt.Println("Key deleted.")

	// Try to retrieve the deleted key
	value, exists := config.GedisClient.Get("key")
	if !exists {
		fmt.Println("Key does not exist.")
	} else {
		fmt.Println("Retrieved value:", value)
	}

	// Rename a key
	config.GedisClient.Set("original", "value")
	err := config.GedisClient.RENAME("original", "renamed")
	if err != nil {
		fmt.Println("RENAME failed:", err)
	} else {
		fmt.Println("Key renamed successfully.")
	}

	// Retrieve the renamed key
	value, exists = config.GedisClient.Get("renamed")
	if !exists {
		fmt.Println("Renamed key does not exist.")
	} else {
		fmt.Println("Retrieved renamed value:", value)
	}

	// List all keys
	keys := config.GedisClient.Keys()
	fmt.Println("All keys in the database:")
	for _, key := range keys {
		fmt.Println(key)
	}
}
