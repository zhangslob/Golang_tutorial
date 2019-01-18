package main

import (
	"fmt"
	"math"
)

type Method struct {
	X, Y float64
}

func (v Method) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Method{3, 4}
	fmt.Println(v)
	fmt.Println(v.Abs())
}
