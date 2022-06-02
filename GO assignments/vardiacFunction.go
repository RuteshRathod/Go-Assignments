package main

import "fmt"

// vardiacFunction is a function that takes a variable number of arguments
// and returns the largest value.
// HERE  PARAMTERS ARE TREADED AS SLICE OF INTEGERS
func vardiacFunction(value ...int) int {
	var max = value[0]
	for _, v := range value {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	// pass in a list of values
	fmt.Println(vardiacFunction(1, 2, 3, 4, 5))
	vlsice := []int{11, 23, 13, 4, 5}
	//Passong slice as argument
	fmt.Println(vardiacFunction(vlsice...))
}
