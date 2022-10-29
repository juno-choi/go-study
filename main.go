package main

import (
	"fmt"
	"strings"
)

func main() {
	length, upper := lenAndUpper("juno")
	fmt.Println("length =", length)
	fmt.Println("upper =", upper)
}

func lenAndUpper(name string) (length int, upper string) {
	defer fmt.Println("lenAndUpper end !")
	fmt.Println("lenAndUpper start !")
	length = len(name)
	upper = strings.ToUpper(name)
	return
}
