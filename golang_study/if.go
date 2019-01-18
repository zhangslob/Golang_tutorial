package main

import (
	"fmt"
	"math"
)

func sari(x float64) string {
	if x < 0 {
		return sari(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sari(2), sari(-4))
}
