package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f\n", e)
}

const e = 1e-8

func Sarto(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	for {
		new_z := z - ((z*z - x) / (2*z))
		if math.Abs(new_z - z) < e {
			return new_z, nil
		}
		z = new_z
	}
}

func main() {
	fmt.Println(Sarto(2))
	fmt.Println(Sarto(-2))
}
