package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	val := make([][]uint8, dy, dy)
	for y, _ := range val {
		elem := make([]uint8, dx, dx)
		for x, _ := range elem {
			elem[x] = uint8(x ^ y)
		}
		val[y] = elem
	}
	return val
}

func main() {
	pic.Show(Pic)
}
