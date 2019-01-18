package main

import "fmt"

var pews = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pews {
		fmt.Println(i, v)
	}
}
