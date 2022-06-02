package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	println("Enter a string 1:")
	fmt.Scan(&s)
	s = strings.ToLower(s)
	check(s)
}
func check(s string) {
	var starts string = "i"
	var end string = "n"
	var contain string = "a"
	var found bool = strings.HasPrefix(s, starts) &&
		strings.HasSuffix(s, end) &&
		strings.Contains(s, contain)

	if found == true {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
}
