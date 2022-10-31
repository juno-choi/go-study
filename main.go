package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	hoby []string
}

func main() {
	hoby := []string{"sing", "golf"}
	juno := person{name: "juno", age: 20, hoby: hoby}
	fmt.Println(juno)
}
