package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)
	var curr string

	for {
		fmt.Printf("Please enter an integer or enter \"X\" to exit: ")
		var _, scan_err = fmt.Scan(&curr)
		var digit, conversion_err = strconv.Atoi(curr)

		switch {
		case scan_err != nil:
			fmt.Print(scan_err)
		case curr == "X":
			return
		case conversion_err == nil:
			sli = append(sli, digit)
			sort.Ints(sli)
			fmt.Println("The sorted array is: ", sli)
		default:
			fmt.Println("Value entered must be either an integer or \"X\" to exit.")
			return

		}
	}
}
