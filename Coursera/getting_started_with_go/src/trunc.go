package main

import "fmt"

func main() {
	var float_value float64
	fmt.Printf("Please enter a floating point number and press enter: ")
	var _, err = fmt.Scan(&float_value)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		int_value := int(float_value)
		fmt.Printf("%d", int_value)
	}
}
