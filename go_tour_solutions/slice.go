package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ys := make([][]uint8, dy)
	for i := range ys {
		ys[i] = make([]uint8, dx)
	}

	for i, xs := range ys {
		for j := range xs {
			xs[j] = uint8((i + j) / 2)
		}
	}
	return ys
}

func main() {
	pic.Show(Pic)
}
