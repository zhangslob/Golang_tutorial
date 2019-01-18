package main

import "fmt"

type Vortexes struct {
	X int
	Y int
}

func main() {
	v := Vortexes{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
