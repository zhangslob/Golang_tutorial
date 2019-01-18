package main

import "fmt"

func main() {
	var s [] int
	printSli(s)

	// 当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。
	s = append(s, 0)
	printSli(s)

	s = append(s, 1)
	printSli(s)

	s = append(s, 2, 3, 4)
	printSli(s)
}

func printSli(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}