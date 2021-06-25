package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Print("Enter filename: ")
	fmt.Scan(&filename)

	people := make([]Person, 0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			names := strings.Fields(scanner.Text())
			if len(names) > 2 {
				fmt.Printf("Too many values in the line \"%s\"", strings.Join(names, " "))
				return
			}
			person := Person{fname: names[0], lname: names[1]}
			people = append(people, person)
		}
	}
	file.Close()

	for _, person := range people {
		fmt.Println(person.fname, person.lname)
	}
}
