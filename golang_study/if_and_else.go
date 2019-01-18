package main

import (
	"fmt"
	"math"
)

func powB(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		powB(3, 2, 10),
		powB(3, 3, 20))
}
