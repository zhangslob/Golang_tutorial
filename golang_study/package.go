package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(2)
	fmt.Println("My favorite number is", rand.Int())
	fmt.Println("Random number:")
}
