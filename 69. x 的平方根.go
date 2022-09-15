package main

import "math"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	x0 := float64(x)
	C := float64(x)
	for {
		x1 := (x0 + C/x0) / 2
		if math.Abs(x0-x1) < 0.1 {
			break
		}
		x0 = x1
	}
	return int(x0)
}
