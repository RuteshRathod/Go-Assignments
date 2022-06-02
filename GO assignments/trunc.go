package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	println("Enter a string:")
	fmt.Scan(&s)

	if strings.Contains(s, "a") {
		if strings.Contains(s, "i") {
			if strings.Contains(s, "n") {
				println("Found")
			} else {
				fmt.Println("Not Found")
			}

		} else {
			fmt.Println("Not Found")
		}
	} else {
		fmt.Println("Not Found")
	}

}
