package operations

import (
	"Gedis-Client/config"
	"fmt"
)

// ListOperations demonstrates list operations using the Gedis client.
func ListOperations() {

	// LPUSH operation

	len, err := config.GedisClient.LPush("mylist", "value1", "value2", "value3")
	if err != nil {
		fmt.Println("LPUSH failed:", err)
	} else {
		fmt.Println("LPUSH successful, new length:", len)
	}

	fmt.Println("List after LPUSH:")
	values, _ := config.GedisClient.LRange("mylist", 0, -1)
	fmt.Println(values)

	// RPUSH operation
	len, err = config.GedisClient.RPush("mylist", "value4", "value5")
	if err != nil {
		fmt.Println("RPUSH failed:", err)
	} else {
		fmt.Println("RPUSH successful, new length:", len)
	}

	fmt.Println("List after RPUSH:")
	values, _ = config.GedisClient.LRange("mylist", 0, -1)
	fmt.Println(values)

	// LRANGE operation
	start, stop := 0, -1
	values, err = config.GedisClient.LRange("mylist", start, stop)
	if err != nil {
		fmt.Println("LRANGE failed:", err)
	} else {
		fmt.Println("LRANGE successful, values:", values)
	}

	// LPOP operation
	value, err := config.GedisClient.LPop("mylist")
	if err != nil {
		fmt.Println("LPOP failed:", err)
	} else {
		fmt.Println("LPOP successful, value:", value)
	}

	fmt.Println("List after LPOP:")
	values, _ = config.GedisClient.LRange("mylist", 0, -1)
	fmt.Println(values)

	// RPOP operation
	value, err = config.GedisClient.RPop("mylist")
	if err != nil {
		fmt.Println("RPOP failed:", err)
	} else {
		fmt.Println("RPOP successful, value:", value)
	}

	fmt.Println("List after RPOP:")
	values, _ = config.GedisClient.LRange("mylist", 0, -1)
	fmt.Println(values)

}
