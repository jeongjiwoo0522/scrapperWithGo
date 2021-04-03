package main

import (
	"fmt"
	"github/jeongjiwoo0522/scrapperWithGo/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	word := "hello"
	dictionary.Add(word, "First")
	dictionary.Update(word, "Second")
	dictionary.Delete(word)
	definition, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}