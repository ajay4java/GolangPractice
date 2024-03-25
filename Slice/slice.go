package main

import "fmt"

func main() {

	myString := "Hello World"
	hello := myString[0:5]
	world := myString[6:1]

	fmt.Println(hello, world)
}
