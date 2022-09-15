package main

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	m := pow(x, n/2)
	if n%2 == 0 {
		return m * m
	} else {
		return m * m * x
	}
}
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	flag := true
	if n < 0 {
		flag = false
		n = -n
	}
	if flag {
		return pow(x, n)
	} else {
		return 1 / pow(x, n)
	}
}
