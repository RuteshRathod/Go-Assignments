package main

import "fmt"

func main() {
	i := 5
	defer fmt.Println(i + 1)
	// Here valye of i shoule be 7 but it prints 6, because value of arguments are evaluated immediately for defer, as soon its defined
	i++
	defer fmt.Println("Bye")
	fmt.Println("Helle world")
}
