package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() string {
	fmt.Print("Enter a list of space separated integers: ")

	reader := bufio.NewReader(os.Stdin)
	values, _ := reader.ReadString('\n')

	return strings.TrimSpace(values)
}

func ConvertToIntArray(input string) ([]int, error) {
	values := make([]int, 0, 10)
	for _, value := range strings.Split(input, " ") {
		converted, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		values = append(values, converted)
	}
	return values, nil
}

func BubbleSort(values []int) {
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(values)-1; i++ {
			if values[i] > values[i+1] {
				Swap(values, i)
				changed = true
			}
		}
	}
}

func Swap(values []int, index int) {
	temp := values[index]
	values[index] = values[index+1]
	values[index+1] = temp
}

func main() {
	input := GetInput()

	values, err := ConvertToIntArray(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	BubbleSort(values)
	fmt.Println("The sorted array is:", values)
}
