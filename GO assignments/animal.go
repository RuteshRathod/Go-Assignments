package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Generic Animal Struct
type Animal struct {
	food, locomotion, noise string
}

// receiver of eat
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

//receiver of move
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

//receiver of speak
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var animal Animal

	for {
		// enter animal name space and enter info you want
		fmt.Println(">Please Enter Details :> Name of Animal & Info >")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		details := scanner.Text()

		// Spliting name and info
		name := strings.Split(details, " ")[0]
		info := strings.Split(details, " ")[1]

		switch name {
		case "cow":
			animal = Animal{"grass", "walks", "moo"}
		case "bird":
			animal = Animal{"worms", "fly", "peep"}
		case "snake":
			animal = Animal{"mice", "slither", "hsss"}
		default:
			fmt.Printf("%s does not occur in (cow, bird, snake)", name)
			return
		}

		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("%s does not occur in (cow, bird, snake)", info)
			return
		}

	}
}
