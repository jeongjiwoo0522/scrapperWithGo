package main

import (
	"fmt"
	"github/jeongjiwoo0522/scrapperWithGo/something"
)

func main() {
	fmt.Println("Hello World")
	something.SayHello()

	var name string = "Cat" //variable
	const age int = 12 // constant
	var id, address string
	var (
		height int
		weight float32
	)
	short := "hello"
	fmt.Println(name, id, address, height, weight, short)
	// variables and constants
}
