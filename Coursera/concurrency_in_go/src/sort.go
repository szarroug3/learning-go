package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetInput() string {
	fmt.Print("Enter space separated integers: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ConvertToIntArray(input string) ([]int, error) {
	values := make([]int, 0)

	for _, value := range strings.Split(input, " ") {
		converted, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		values = append(values, converted)
	}

	return values, nil
}

func Split(sli []int, count int) [][]int {
	increment := int(len(sli) / count)
	slices := make([][]int, 0)

	start := 0
	end := increment
	for index := 0; index < count; index++ {
		if index == count-1 {
			end = len(sli)
		}
		slices = append(slices, sli[start:end])
		start = start + increment
		end = end + increment
	}

	return slices
}

func Sort(sli []int, channel chan []int) {
	fmt.Println("Sorting:", sli)
	sort.Ints(sli)
	channel <- sli
}

func Merge(count int, channel chan []int) []int {
	sorted_slices := make([][]int, 0)

	for index := 0; index < count; index++ {
		sorted_slices = append(sorted_slices, <-channel)
	}

	result := make([]int, 0)
	var min int
	var index int
	found := true

	for found {
		index = -1
		found = false
		for curr_index, sli := range sorted_slices {
			if len(sli) == 0 {
				continue
			}

			found = true
			if index == -1 {
				min = sli[0]
				index = curr_index
			} else if sli[0] < min {
				min = sli[0]
				index = curr_index
			}
		}

		if found {
			result = append(result, sorted_slices[index][0])
			sorted_slices[index] = sorted_slices[index][1:]
		}
	}
	return result
}

func main() {
	chunk_count := 4
	input := GetInput()

	values, err := ConvertToIntArray(input)
	if err != nil {
		fmt.Println(err)
	}

	slices := Split(values, chunk_count)
	channel := make(chan []int, chunk_count)

	for _, sli := range slices {
		go Sort(sli, channel)
	}

	sorted := Merge(4, channel)
	fmt.Println("Sorted:", sorted)
}
