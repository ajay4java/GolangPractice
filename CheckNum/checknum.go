package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	str1, _ := reader.ReadString('\n')

	// remove newline
	str1 = strings.TrimSpace(str1)

	// convert string variable to int variable
	num, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}

	fmt.Println(num >= 1 && num <= 10)
}
