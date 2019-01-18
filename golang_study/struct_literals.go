package main

import "fmt"

type Some struct {
	X, Y int
}

var (
	v1 = Some{1, 2}
	v2 = Some{X: 1}
	v3 = Some{}
	p = &Some{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
