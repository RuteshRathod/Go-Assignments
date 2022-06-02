package main

import (
	"fmt"
	"sort"
	"strconv"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func bubbleSort(arr []int) {
	sort.Ints(arr)
}

func main() {
	var arr []int
	var scanInt string
	fmt.Println("Enter Integers and press 'e' to exit:")
	// Scanner for loop to read input
	for {
		_, err := fmt.Scan(&scanInt)
		if scanInt == "e" {
			fmt.Println("Exiting Loop", err)
			break
		}
		// Convert string to int
		convertedSlice, err := strconv.Atoi(scanInt)
		// Append to slice
		arr = append(arr, convertedSlice)
	}
	bubbleSort(arr)
	fmt.Println(arr)
}
