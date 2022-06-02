package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice1 := make([]int, 0, 3)

	var scanSlice string

	fmt.Println("Enter the slice elements and Press 'e' to exit: ")

	for {
		fmt.Println("Enter an Integer: ")
		_, err := fmt.Scan(&scanSlice)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if scanSlice == "e" {
			fmt.Println("Exiting the loop")
			break
		}
		convertedInt, err := strconv.Atoi(scanSlice)
		slice1 = append(slice1, convertedInt)
		sort.Ints(slice1)
		fmt.Println("Slice: ", slice1)
	}
}
