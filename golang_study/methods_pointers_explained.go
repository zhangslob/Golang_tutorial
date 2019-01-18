package main

import (
	"fmt"
	"math"
)

type SecondVertex struct {
	X, Y float64
}

func AbsSecond(v SecondVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *SecondVertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := SecondVertex{3, 4}
	Scale(&v, 10)
	fmt.Println(AbsSecond(v))
}
