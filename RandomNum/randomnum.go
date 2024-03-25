package main

import (
	"fmt"
	"math/rand"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	randomNum := random(1, 7)
	fmt.Printf("Random number: %d\n", randomNum)

}
