package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func checkForValueInSlice(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func checkIfMapContains(inputMap map[string]Animal, valueToBeChecked string) bool {
	if _, ok := inputMap[valueToBeChecked]; ok {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	informationType := []string{"eat", "move", "speak"}
	animalNames := []string{"cow", "bird", "snake"}
	animalNameTypeMap := make(map[string]Animal)

	fmt.Println("Only two commands are supported and the command should be requested using below format.")
	fmt.Println("1. newanimal animal_name animal_type")
	fmt.Println("animal_type can be only one of the cow, bird or snake")
	fmt.Println("2. query animal_name info_type")
	fmt.Println("info_type can be only one of the eat, move or speak")
	fmt.Println("Each command should have exactly three arguments and mentioned above.")
	fmt.Println("Enter x to exit.")

	for {
		// Read inputs
		fmt.Print("> ")
		inputStr, _ := reader.ReadString('\n')
		inputStr = strings.Trim(inputStr, "\n")
		inputStr = strings.ToLower(inputStr)

		// Exit the program
		if inputStr == "x" {
			fmt.Println("Exiting the program.")
			os.Exit(0)
		}

		inputs := strings.Split(inputStr, " ")

		// Check if the input provided contains 3 strings
		if len(inputs) > 3 || len(inputs) < 3 {
			fmt.Println("Enter exactly 3 values in the input.")
			continue
		}

		// Check if the first value is the known command value
		if inputs[0] != "newanimal" && inputs[0] != "query" {
			fmt.Println("The first value of the command can either be newanimal or query")
			continue
		}

		// Check if the third value belong to the predefined values
		if inputs[0] == "newanimal" && !checkForValueInSlice(animalNames, inputs[2]) {
			fmt.Println("For a newanimal command, the third argument should be one of the values => cow, bird, snake")
			continue
		} else if inputs[0] == "query" && !checkForValueInSlice(informationType, inputs[2]) {
			fmt.Println("For a query command, the third argument should be one of the values => eat, move, speak")
			continue
		}

		// Process newanimal command
		if inputs[0] == "newanimal" {
			// Check if the animal is already created
			if checkIfMapContains(animalNameTypeMap, inputs[1]) {
				fmt.Println("Animal with the name " + inputs[1] + " already created, use any other name.")
				continue
			}

			// Create an entry into the map
			if inputs[2] == "cow" {
				animalNameTypeMap[inputs[1]] = Cow{name: inputs[1]}
			} else if inputs[2] == "bird" {
				animalNameTypeMap[inputs[1]] = Bird{name: inputs[1]}
			} else if inputs[2] == "snake" {
				animalNameTypeMap[inputs[1]] = Snake{name: inputs[1]}
			}

			fmt.Println("Created it!")
		} else if inputs[0] == "query" {
			// Check if the animal whose information is requested is already present
			if !checkIfMapContains(animalNameTypeMap, inputs[1]) {
				fmt.Println("Animal with the name " + inputs[1] + " not present, please create a new one with the newanimal command.")
				continue
			}

			// Print information
			if inputs[2] == "eat" {
				animalNameTypeMap[inputs[1]].Eat()
			} else if inputs[2] == "move" {
				animalNameTypeMap[inputs[1]].Move()
			} else if inputs[2] == "speak" {
				animalNameTypeMap[inputs[1]].Speak()
			}
		}
	}
}
