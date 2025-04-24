package operations

/*
package tests

import (
	"testing"
	"time"

	"github.com/GedisCaching/Gedis/gedis"
)

func TestTTLOperations(t *testing.T) {
	// Create a new Gedis instance with default config
	g, err := gedis.NewGedis(gedis.Config{})
	if err != nil {
		t.Fatalf("Failed to create Gedis instance: %v", err)
	}

	// Test SetWithExpiry and TTL
	t.Run("SetWithExpiry and TTL", func(t *testing.T) {
		// Set key with 1 second expiry
		g.SetWithExpiry("expiring", "value", 1*time.Second)

		// Check TTL immediately
		ttl, exists := g.TTL("expiring")
		if !exists {
			t.Error("TTL failed: key not found")
		}
		if ttl <= 0 {
			t.Errorf("Expected positive TTL, got %v", ttl)
		}

		// Check that key exists
		val, exists := g.Get("expiring")
		if !exists {
			t.Error("Key should exist immediately after setting")
		}
		if val != "value" {
			t.Errorf("Expected 'value', got %v", val)
		}

		// Wait for expiration
		time.Sleep(1100 * time.Millisecond)

		// Check that key has expired
		_, exists = g.Get("expiring")
		if exists {
			t.Error("Key should have expired")
		}
	})

	// Test DEXPIRE
	t.Run("DEXPIRE Operation", func(t *testing.T) {
		// Set a normal key without expiry
		g.Set("key1", "value1")

		// Add expiry
		err := g.DEXPIRE("key1", 1*time.Second)
		if err != nil {
			t.Errorf("DEXPIRE failed: %v", err)
		}

		// Check TTL
		ttl, exists := g.TTL("key1")
		if !exists {
			t.Error("TTL failed: key not found after DEXPIRE")
		}
		if ttl <= 0 {
			t.Errorf("Expected positive TTL after DEXPIRE, got %v", ttl)
		}

		// Wait for expiration
		time.Sleep(1100 * time.Millisecond)

		// Check that key has expired
		_, exists = g.Get("key1")
		if exists {
			t.Error("Key should have expired after DEXPIRE")
		}
	})
}

*/

import (
	"Gedis-Client/config"
	"fmt"
	"time"
)

// TTLOperations demonstrates TTL operations using the Gedis client.
func TTLOperations() {
	// Example usage of the Gedis client
	config.GedisClient.Set("key", "value")

	// Set a key with an expiry time of 5 seconds
	config.GedisClient.SetWithExpiry("expiring_key", "value", 5*time.Minute)

	// Retrieve the TTL of the key
	ttl, exists := config.GedisClient.TTL("expiring_key")
	if exists {
		fmt.Println("TTL of expiring_key:", ttl)
	} else {
		fmt.Println("expiring_key does not exist.")
	}

	// Wait for the key to expire
	config.GedisClient.TTL("expiring_key")

	// Check if the key still exists after expiry
	value, exists := config.GedisClient.Get("expiring_key")
	if exists {
		fmt.Println("expiring_key still exists with value:", value)
	} else {
		fmt.Println("expiring_key has expired.")
	}

	// Set a key with a specific expiry time
	config.GedisClient.SetWithExpiry("temporary_key", "temp_value", 2*time.Second)

	time.Sleep(3 * time.Second)

	// Check if the key still exists after expiry
	value, exists = config.GedisClient.Get("temporary_key")
	if exists {
		fmt.Println("temporary_key still exists with value:", value)
	} else {
		fmt.Println("temporary_key has expired.")
	}
}
