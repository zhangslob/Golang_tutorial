package main

import (
	"fmt"
	"math"
)

//func Sqrt(x float64) float64 {
//	z := float64(2.)
//	s := float64(0)
//	for i:=0; i<10; i++ {
//		z = z -(z*z -x)/(2*z)
//		fmt.Println(i, z)
//		if math.Abs(z - s) < 1e-10 {
//			break
//		}
//		s = z
//	}
//	return z
//}

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		if math.Abs(z-(z-(z*z-x)/(z*2))) < 0.00000000000001 {
			return z
		} else {
			z = z - (z*z-x)/(z*2)
			fmt.Println(">>>>>z is:", z)
		}
	}
}


func main() {
	fmt.Println(Sqrt(2))
}
