package main

import (
	"fmt"
)

func main() {
	array("1", "2", "3", "4", "5")
}

func array(name ...string) {
	fmt.Println(name)
}
