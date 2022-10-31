package main

import (
	"fmt"

	"github.com/project/helloworld/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "hello~")
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
	dictionary.Delete(baseWord)
	word2, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word2)
}
