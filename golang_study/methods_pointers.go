package main

import (
	"fmt"
	"math"
)

type MyVertex struct {
	X, Y float64
}

func (v MyVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v MyVertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X, v.Y)
}

func main() {
	v := MyVertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
