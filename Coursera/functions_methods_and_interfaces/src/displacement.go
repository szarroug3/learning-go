package main

import "fmt"

func GetInput(name string) (float64, error) {
	var value float64

	fmt.Printf("Enter %s: ", name)
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		result := .5 * acceleration * time * time
		result = result + (velocity * time)
		result = result + displacement
		return result
	}
	return fn
}

func main() {
	acceleration, err := GetInput("acceleration")
	if err != nil {
		fmt.Print(err)
		return
	}
	velocity, err := GetInput("velocity")
	if err != nil {
		fmt.Print(err)
		return
	}
	displacement, err := GetInput("displacement")
	if err != nil {
		fmt.Print(err)
		return
	}
	time, err := GetInput("time")
	if err != nil {
		fmt.Print(err)
		return
	}

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println(fn(time))
}
