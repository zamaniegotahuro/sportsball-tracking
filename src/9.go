  package main

import (
"fmt"
"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	fmt.Println(randomNumber)
}