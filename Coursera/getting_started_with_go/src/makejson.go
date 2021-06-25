package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	person := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	person["name"] = strings.TrimSpace(name)

	fmt.Print("Enter address: ")
	address, _ := reader.ReadString('\n')
	person["address"] = strings.TrimSpace(address)

	json_person, _ := json.Marshal(person)
	fmt.Println(string(json_person))
}
