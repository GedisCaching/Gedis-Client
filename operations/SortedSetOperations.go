package operations

import (
	"Gedis-Client/config"
	"fmt"
)

// SortedSetOperations demonstrates sorted set operations using the Gedis client.

func SortedSetOperations() {
	// Example usage of the Gedis client
	config.GedisClient.ZAdd("myset", map[string]float64{"member1": 1.0, "member2": 2.0})

	// Retrieve the members in the sorted set
	members := config.GedisClient.ZRange("myset", 0, -1, false)
	fmt.Println("Members in sorted set:", members)

	// Get the rank of a member
	rank, exists := config.GedisClient.ZRank("myset", "member1")
	if exists {
		fmt.Println("Rank of member1:", rank)
	} else {
		fmt.Println("member1 does not exist in the sorted set.")
	}

	// Get the rank of a non-existent member
	rank, exists = config.GedisClient.ZRank("myset", "nonexistent")
	if exists {
		fmt.Println("Rank of nonexistent member:", rank)
	} else {
		fmt.Println("nonexistent member does not exist in the sorted set.")
	}

}
