package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a)
	printSlices("a", a)

	b := make([]int, 0, 5)
	fmt.Println(b)
	printSlices("b", b)

	c := b[:2]
	fmt.Println(c)
	printSlices("c", c)

	d := c[2:5]
	fmt.Println(d)
	printSlices("d", d)
}

func printSlices(s string, x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}