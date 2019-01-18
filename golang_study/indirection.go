package main

import "fmt"

type ThirdVertex struct {
	X, Y float64
}

func (v *ThirdVertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *ThirdVertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := ThirdVertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &ThirdVertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
