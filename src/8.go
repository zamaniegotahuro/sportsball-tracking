// Package main implements a sportsball tracking web application.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate a random number between 1 and 9
	randNum := rand.Intn(9) + 1

	// Print the random number
	fmt.Println("The random number is", randNum)
}
