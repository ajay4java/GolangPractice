package main

import "fmt"

func main() {

	var a float64 = 3
	var b float64 = 6
	c := Sum(a, b)

	fmt.Println(c)
}

func Sum(a, b float64) float64 {
	return a + b
}
