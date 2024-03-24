package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v", name)
}
