package main

import (
	"fmt"
	"time"
)

func main() {
	var cache1 = newCache()
	var ch, value int
	var ttl int64
	var key string
	go checkExpiry(&cache1)
	for {
		fmt.Println(time.Now().Unix())
		fmt.Println("Enter choice: ")
		fmt.Println("1. Set Value with default TTl \n2. Set Value with TTl \n 3. Get Value \n 4. Delete Value \n 5. Print")
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			fmt.Print("Enter key value pair: ")
			fmt.Scan(&key, &value)
			cache1.set(Parameters{data: value, key: key})
		case 2:
			fmt.Print("Enter key, value and expiration time: ")
			fmt.Scan(&key, &value, &ttl)
			cache1.set(Parameters{data: value, key: key, ttl: ttl})
		case 3:
			fmt.Print("Enter key to get value: ")
			fmt.Scan(&key)
			fmt.Println(cache1.get(key))
		case 4:
			fmt.Print("Enter key to be deleted: ")
			fmt.Scan(&key)
			cache1.delete(key)
		case 5:
			fmt.Print(cache1.Map)
		default:
			break
		}
	}
}
