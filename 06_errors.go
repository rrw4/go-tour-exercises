package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		z := 1.0
		diff := .0000001
		last_z := z
		for {
			z = last_z - ((last_z*last_z - x) / (2 * last_z))
			dist := z - last_z
			if dist < 0 {
				dist = 0 - dist
			}
			if dist < diff {
				return z, nil
			}
			last_z = z
		}
	}
	//not sure on what to return for the float64 value if error is present
	return 0, ErrNegativeSqrt(-2)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
