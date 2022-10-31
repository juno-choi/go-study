package main

import (
	"fmt"
)

func main() {
	names := map[string]string{"key": "value", "key2": "value2", "key3": "value3"}
	for key, value := range names {
		fmt.Println(key, value)
	}
}
