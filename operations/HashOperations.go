package operations

import (
	"Gedis-Client/config"
	"fmt"
)

// HashOperations demonstrates hash operations using the Gedis client.
func HashOperations() {
	// Example usage of the Gedis client
	config.GedisClient.HSET("user:1", "name", "Gerges Hany")
	config.GedisClient.HSET("user:1", "email", "gerges@example.com")
	config.GedisClient.HSET("user:1", "age", 30)
	config.GedisClient.HSET("user:1", "address", "123 Main St")
	config.GedisClient.HSET("user:1", "phone", "123-456-7890")
	config.GedisClient.HSET("user:1", "occupation", "Software Engineer")
	config.GedisClient.HSET("user:1", "company", "Tech Corp")

	// Retrieve the fields in the hash
	fields, exists := config.GedisClient.HKEYS("user:1")
	if exists {
		fmt.Println("Fields in hash user:1:", fields)
	} else {
		fmt.Println("Hash user:1 does not exist.")
	}

	// Retrieve the value of a specific field
	name, exists := config.GedisClient.HGET("user:1", "name")
	if exists {
		fmt.Println("Name of user:1:", name)
	} else {
		fmt.Println("Field name does not exist in hash user:1.")
	}

	// Retrieve all fields and values in the hash
	userData, exists := config.GedisClient.HGETALL("user:1")
	if exists {
		fmt.Println("All fields and values in hash user:1:", userData)
	} else {
		fmt.Println("Hash user:1 does not exist.")
	}

	// Delete a field from the hash
	deleted, err := config.GedisClient.HDEL("user:1", "email")
	if deleted {
		fmt.Println("Field email deleted from hash user:1.")
	} else {
		fmt.Println("Failed to delete field email from hash user:1:", err)
	}

	// Check the length of the hash
	length, exists := config.GedisClient.HLEN("user:1")
	if exists {
		fmt.Println("Length of hash user:1:", length)
	} else {
		fmt.Println("Hash user:1 does not exist.")
	}

}
