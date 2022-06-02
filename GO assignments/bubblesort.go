// Dynamic array sort using bubble sort,
// give inputs of negative numbers and positive numbers
// and it will sort the array in ascending order
// linkedIn: https://www.linkedin.com/in/ruteshrathod/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Slice of ints
	arrSlice := make([]int, 0, 3)
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
		arrSlice = append(arrSlice, convertedSlice)
	}
	// Calling the function BubbleSort via passing appended slice
	BubbleSort(arrSlice)
	// Printing the sorted array
	fmt.Println("Sorted Array:", arrSlice)

}

// BubbleSort Function to sort the array
func BubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a

}
