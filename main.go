package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a
	*b = 20
	fmt.Println(a, *b)
}
