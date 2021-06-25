/*
 * This file contains two functions, IncrementX and PrintX. If these functions
 * were to be run in parallel, they would cause a race condition because they
 * use the variable "x" and the output of the PrintX function would be
 * dependent on if it was to excecute the print statement before or after
 * IncrementX's incrememnt statement. Running this programming repeatedly
 * shows that the output you get can be different each time. Sometimes, the
 * same number will be printed multiple times. Other times, numbers may be
 * skipped. This program cannot guarantee that the numbers will be printed
 * 1-100000 as expected.
 */

package main

import (
	"fmt"
	"time"
)

func IncrementX(x *int) {
	fmt.Println("Starting IncrementX")
	for i := 0; i < 10000; i++ {
		*x++
		time.Sleep(time.Second)
	}
	fmt.Println("Ending IncrementX")
}

func PrintX(x *int) {
	fmt.Println("Starting PrintX")
	for *x < 10000 {
		fmt.Println(*x)
		time.Sleep(time.Second)
	}
	fmt.Println("Ending PrintX")
}

func main() {
	x := 1
	go IncrementX(&x)
	go PrintX(&x)

	go func(msg string) {
		fmt.Println(msg)
	}("Starting goroutines")

	time.Sleep(time.Minute)
}
