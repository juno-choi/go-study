package main

import (
	"fmt"
)

func main() {
	fmt.Println(zeroOrOne(1))
}

func zeroOrOne(number int) string {
	switch badNumber := number + 2; badNumber {
	case 0:
		return "zero"
	case 1:
		return "one"
	}
	return "bad request"
}
