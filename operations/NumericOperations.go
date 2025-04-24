package operations

import (
	"Gedis-Client/config"
	"fmt"
)

// NumericOperations demonstrates numeric operations using the Gedis client.
func NumericOperations() {
	// Example usage of the Gedis client
	config.GedisClient.Set("counter", 0)

	// Increment the counter
	newValue, err := config.GedisClient.Incr("counter")
	if err != nil {
		fmt.Println("INCR failed:", err)
	} else {
		fmt.Println("Counter after INCR:", newValue)
	}

	// Decrement the counter
	newValue, err = config.GedisClient.Decr("counter")
	if err != nil {
		fmt.Println("DECR failed:", err)
	} else {
		fmt.Println("Counter after DECR:", newValue)
	}
}
