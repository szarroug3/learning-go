package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func GetInput() (string, string, string, error) {
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	values := strings.Split(input, " ")
	if len(values) != 3 {
		return "", "", "", errors.New("Input must have exactly three values.")
	}

	return values[0], values[1], values[2], nil
}

func GetCommandFunc(command string) (func(map[string]Animal, string, string) error, error) {
	switch strings.ToLower(command) {
	case "newanimal":
		return CreateNewAnimal, nil
	case "query":
		return QueryAnimal, nil
	default:
		return nil, errors.New(fmt.Sprintf("Command \"%s\" could not be found.", command))
	}
}

func CreateNewAnimal(animals map[string]Animal, name string, animal_type string) error {
	_, ok := animals[name]
	if ok {
		return errors.New(fmt.Sprintf("Animal with name \"%s\" already exists.", name))
	}

	switch strings.ToLower(animal_type) {
	case "cow":
		animals[name] = Cow{name}
	case "bird":
		animals[name] = Bird{name}
	case "snake":
		animals[name] = Snake{name}
	default:
		return errors.New(fmt.Sprintf("Type \"%s\" could not be found.", animal_type))
	}

	fmt.Println("Created it!")
	return nil
}

func QueryAnimal(animals map[string]Animal, name string, request string) error {
	animal, ok := animals[name]
	if !ok {
		return errors.New(fmt.Sprintf("Animal with name \"%s\" cannot be found.", name))
	}

	switch strings.ToLower(request) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		return errors.New(fmt.Sprintf("Request \"%s\" could not be found.", request))
	}

	return nil
}

func main() {
	animals := make(map[string]Animal)

	for {
		command, name, value, err := GetInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fn, err := GetCommandFunc(command)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = fn(animals, name, value)
		if err != nil {
			fmt.Println(err)
		}
	}
}
