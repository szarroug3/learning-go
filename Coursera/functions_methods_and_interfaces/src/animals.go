package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func GetInput() (string, string, error) {
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	values := strings.Split(input, " ")
	if len(values) != 2 {
		return "", "", errors.New("Input must have exactly two values.")
	}

	return values[0], values[1], nil
}

func GetAnimal(animals map[string]Animal, name string) (Animal, error) {
	if animal, ok := animals[strings.ToLower(name)]; ok {
		return animal, nil
	}
	return Animal{}, errors.New(fmt.Sprintf("Animal \"%s\" could not be found.", name))
}

func GetAnimalFunc(animal Animal, request string) (func(), error) {
	switch strings.ToLower(request) {
	case "eat":
		return animal.Eat, nil
	case "move":
		return animal.Move, nil
	case "speak":
		return animal.Speak, nil
	default:
		return nil, errors.New(fmt.Sprintf("Request \"%s\" could not be found.", request))
	}
}

func main() {
	animals := make(map[string]Animal)

	animals["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		name, request, err := GetInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		animal, err := GetAnimal(animals, name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		animal_func, err := GetAnimalFunc(animal, request)
		if err != nil {
			fmt.Println(err)
			continue
		}

		animal_func()
	}
}
