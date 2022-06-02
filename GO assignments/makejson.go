package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {

	personMap := make(map[string]string)
	var name string
	var address string

	fmt.Println("Enter Person Name:")
	fmt.Scanln(&name)

	fmt.Println("Enter Person address:")
	fmt.Scanln(&address)

	personMap["name"] = name
	personMap["address"] = address

	jsonString, _ := json.Marshal(personMap)
	fmt.Println(string(jsonString))

}
