package functions

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	z2 := x
	for i:= 0; z2 > 0.001 && i < 10; i++ {
		z -= (z*z - x) / (2 * x)
		z2 = math.Abs(z*z - x)
		fmt.Printf("i: %d  z: %g Abs(z):%g\n", i, z, z2)
	}
	return z
}
