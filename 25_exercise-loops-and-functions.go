package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, int) {
	z := 1.0
	threshold := 1.0e-10
	count := 0
	pre := 0.0
	d := 1.0
	for d > threshold {
		z = z - (z*z-x)/(2.0*z)
		count++
		if d == math.Abs(z-pre) {
			break
		} else {
			d = math.Abs(z - pre)
			pre = z
		}
	}
	return z, count
}

func main() {
	var num float64 = 3
	newton, loop := sqrt(num)
	root := math.Sqrt(num)
	fmt.Printf("Number: %v\nsqrt() [Newton method]:\t%g\nSqrt() [math.Sqrt]:\t%g\nLoop count: %v\n", num, newton, root, loop)
}
