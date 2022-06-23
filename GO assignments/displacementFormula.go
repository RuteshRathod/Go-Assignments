package main

import "fmt"

// GenDisplaceFn generates a function that calculates displacement
func GenDisplaceFn(acceleration, velocity, s0 float64) func(time float64) float64 {
	fn := func(time float64) float64 {
		displacement := (0.5 * acceleration * time * time) + (velocity * time) + s0
		return displacement
	}
	return fn
}

func main() {
	var acceleration, velocity, s0, t float64

	fmt.Print("Enter accelerationeleration: ")
	fmt.Scan(&acceleration)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&velocity)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)

	// Create a function that calculates displacement
	calculateDisplacement := GenDisplaceFn(acceleration, velocity, s0)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	// Call the function and print the result
	fmt.Printf("Displacement at time=%.2f is %f ", t, calculateDisplacement(t))
}
