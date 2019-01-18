package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	fmt.Println(s)
	printSlice(s)

	s = s[:4]
	fmt.Println(s)
	printSlice(s)

	s = s[:2]
	fmt.Println(s)
	printSlice(s)
}

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
