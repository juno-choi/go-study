package main

import (
	"fmt"
)

func main() {
	array(5, 4, 3, 2, 1)
}

func array(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}
