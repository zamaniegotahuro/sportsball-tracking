package main

import "math/rand"

func main() {
	rand.Seed(9)
	code := rand.Intn(10)
	fmt.Println(code)
}
