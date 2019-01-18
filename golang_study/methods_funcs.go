package main

import (
	"fmt"
	"math"
)

type Ver struct {
	X, Y float64
}

func Abs(v Ver) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Ver{3, 6}
	fmt.Println(Abs(v))
}
