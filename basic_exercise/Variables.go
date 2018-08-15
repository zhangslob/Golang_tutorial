//In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	//var ss = "1234"
	//ss declared and not used
	//if you have declared a variables, you must use it !!
}
