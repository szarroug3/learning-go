package main

import (
	"fmt"
	"strings"
)

func main() {
	var result string
	fmt.Print("Please enter a string: ")
	fmt.Scan(&result)
	result = strings.ToLower(result)

	if strings.HasPrefix(result, "i") && strings.Contains(result, "a") && strings.HasSuffix(result, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
