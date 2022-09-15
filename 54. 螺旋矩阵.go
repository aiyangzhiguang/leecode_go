package main

func spiralOrder(matrix [][]int) (ans []int) {
	var f func(int, int, int, int)
	f = func(lr int, lc int, rr int, rc int) {
		if lc > rc || lr > rr {
			return
		}
		if lr == rr && lc == rc {
			ans = append(ans, matrix[lc][rc])
			return
		}
		if lr == rr {
			for i := lc; i <= rc; i++ {
				ans = append(ans, matrix[lr][i])
			}
			return
		}
		if lc == rc {
			for i := lr; i <= rr; i++ {
				ans = append(ans, matrix[i][lc])
			}
			return
		}
		for i := lc; i <= rc; i++ {
			ans = append(ans, matrix[lr][i])
		}
		for i := lr + 1; i <= rr; i++ {
			ans = append(ans, matrix[i][rc])
		}
		for i := rc - 1; i >= lc; i-- {
			ans = append(ans, matrix[rr][i])
		}
		for i := rr - 1; i > lr; i-- {
			ans = append(ans, matrix[i][lc])
		}
		f(lr+1, lc+1, rr-1, rc-1)
	}
	f(0, 0, len(matrix)-1, len(matrix[0])-1)
	return
}
