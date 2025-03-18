package main

import "math/rand"

func main() {
	// Generate a random number between 1 and 6
	num := rand.Intn(6) + 1
	fmt.Println("The number is:", num)
}
