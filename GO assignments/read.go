package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fName string
	lName string
}

func main() {
	// read the string from file
	var filname string
	structSlice := make([]Name, 0)

	fmt.Println("Enter the file name:")
	fmt.Scanln(&filname)
	f, err := os.Open(filname)
	if err != nil {
		fmt.Println(err)
	}
	fileContent := bufio.NewScanner(f)
	for fileContent.Scan() {
		nameText := strings.Split(fileContent.Text(), " ")
		if len(nameText[0]) >= 20 {
			fmt.Println("First Name is too long, taking first 20 characters")
			runes := []rune(nameText[0])
			nameText[0] = string(runes[0:20])
		}
		if len(nameText[1]) >= 20 {
			fmt.Println("Last Name is too long, taking first 20 characters")
			runes := []rune(nameText[1])
			nameText[0] = string(runes[0:20])
		}

		var n1 Name
		n1.fName, n1.lName = nameText[0], nameText[1]
		structSlice = append(structSlice, n1)
	}
	for _, v := range structSlice {
		fmt.Println(v.fName, v.lName)
	}

}
