# Gedis Client

Gedis Client is a Go-based client for interacting with the Gedis caching server, a Redis-like in-memory data store. This client provides a simple and intuitive way to interact with Gedis server, allowing you to perform various operations like basic key-value operations, list operations, hash operations, sorted sets, and more.

## Features

- Basic key-value operations (SET, GET, DELETE)
- Numeric operations (INCR, DECR)
- List operations (LPUSH, RPUSH, LPOP, RPOP, LRANGE)
- Sorted set operations (ZADD, ZRANGE)
- Hash operations (HSET, HGET, HGETALL, HDEL, HKEYS, HLEN)
- TTL operations (EXPIRE, TTL)

## Installation

### Prerequisites

- Go 1.16 or higher
- Gedis server running (locally or remotely)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/GedisCaching/Gedis-Client.git
   cd Gedis-Client
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure the client:
   - Copy the example environment file:
     ```bash
     cp env.example .env
     ```
   - Edit the `.env` file with your Gedis server address and password:
     ```
     ADDRESS=localhost:6379
     PASSWORD=yourPassword
     ```

## Usage

### Running the Example Client

The example client demonstrates various operations supported by Gedis:

```bash
go run main.go
```

### Basic Operations

```go
import (
    "github.com/GedisCaching/Gedis-Client/config"
)

// Initialize client
config.Init()

// Set a key
config.GedisClient.Set("key", "value")

// Get a key
value, exists := config.GedisClient.Get("key")

// Delete a key
config.GedisClient.Delete("key")

// Rename a key
config.GedisClient.RENAME("original", "renamed")

// List all keys
keys := config.GedisClient.Keys()
```

### Numeric Operations

```go
// Increment a value
newValue := config.GedisClient.INCR("counter")

// Decrement a value
newValue := config.GedisClient.DECR("counter")
```

### List Operations

```go
// Add items to the beginning of a list
config.GedisClient.LPUSH("mylist", "item1", "item2")

// Add items to the end of a list
config.GedisClient.RPUSH("mylist", "item3", "item4")

// Remove and get the first item
item, exists := config.GedisClient.LPOP("mylist")

// Remove and get the last item
item, exists := config.GedisClient.RPOP("mylist")

// Get a range of items
items, exists := config.GedisClient.LRANGE("mylist", 0, -1)
```

### Hash Operations

```go
// Set a field in a hash
config.GedisClient.HSET("user:1", "name", "John Doe")

// Get a field from a hash
value, exists := config.GedisClient.HGET("user:1", "name")

// Get all fields and values in a hash
data, exists := config.GedisClient.HGETALL("user:1")

// Delete a field from a hash
deleted, err := config.GedisClient.HDEL("user:1", "name")

// Get all field names in a hash
fields, exists := config.GedisClient.HKEYS("user:1")

// Get the number of fields in a hash
length, exists := config.GedisClient.HLEN("user:1")
```

### TTL Operations

```go
// Set a key with an expiration time (in seconds)
config.GedisClient.EXPIRE("key", 60)

// Get the remaining time to live for a key
ttl, exists := config.GedisClient.TTL("key")
```

### Sorted Set Operations

```go
// Add members to a sorted set
config.GedisClient.ZADD("myzset", 1, "one")

// Get a range of members from a sorted set
members, exists := config.GedisClient.ZRANGE("myzset", 0, -1)
```

## License

This project is licensed under the MIT License.

## Acknowledgments

- [Gedis Caching Server](https://github.com/GedisCaching/Gedis)