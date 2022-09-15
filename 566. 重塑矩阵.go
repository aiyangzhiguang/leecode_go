package main

func matrixReshape(mat [][]int, r int, c int) (ans [][]int) {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	for i := 0; i < r; i++ {
		var tmp []int
		for j := 0; j < c; j++ {
			index := i*c + j
			tmp = append(tmp, mat[index/n][index%n])
		}
		ans = append(ans, tmp)
	}
	return
}
