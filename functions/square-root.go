package functions

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	z := x
	z2 := x
	for i := 0; z2 > 0.001 && i < 10; i++ {
		z -= (z*z - x) / (2 * x)
		z2 = math.Abs(z*z - x)
		fmt.Printf("i: %d  z: %g Abs(z):%g\n", i, z, z2)
	}
	return z, nil
}
