package main

import "fmt"

type Vertexes struct {
	X int
	Y int
}

func main() {
	v := Vertexes{1, 2}
	v.X = 4
	fmt.Println(v)
}
