package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("C://Users//91992//OneDrive//Desktop//file.txt")
	if err != nil {
		return
	}
	defer file.Close()
	var a = []string{"Rio", "Ajay"}
	for i := 0; i < len(a); i++ {
		file.WriteString(a[i])
		file.WriteString("\n")
	}
	data, err := os.ReadFile("C://Users//91992//OneDrive//Desktop//file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))

}
