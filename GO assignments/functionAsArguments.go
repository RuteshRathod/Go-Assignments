package main

import "fmt"

var funcVar func(int) int

func applyIt(afunction func(int) int, val int) int {
	return afunction(val)
}

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }

func main() {
	fmt.Println(applyIt(incFn, 3))
	fmt.Println(applyIt(decFn, 3))
}
