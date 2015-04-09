package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
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
			return z
		}
		last_z = z
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
