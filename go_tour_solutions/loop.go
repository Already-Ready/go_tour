package main

import (
	"fmt"
	"math"
)

const diff = 0.000001

func Sqrt(x float64) float64 {
	z := 1.
	n := 0.
	for math.Abs(n-z) > diff {
		n, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	v := 2.
	v_math, v_close := math.Sqrt(v), Sqrt(v)

	fmt.Printf("value of math : %.8f , value of func : %.8f, diff of two : %.8f", v_math, v_close, v_math-v_close)
}
