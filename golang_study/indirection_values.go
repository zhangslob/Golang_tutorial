package main

import (
	"fmt"
	"math"
)

type FifVertex struct {
	X, Y float64
}

func (v FifVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsF(v FifVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := FifVertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsF(v))

	p := &FifVertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsF(*p))
}
