package main

import "fmt"

func main() {
	primes := [6]int{1, 3, 5, 7, 9, 11}

	var s []int = primes[1:4]
	fmt.Println(s)
}
