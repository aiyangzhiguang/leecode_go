package main

func generate(numRows int) (ans [][]int) {
	ans = append(ans, []int{1})
	for i := 1; i < numRows; i++ {
		var tmp []int
		for j := 0; j < i+1; j++ {
			switch j {
			case 0:
				tmp = append(tmp, ans[i-1][0])
			case i:
				tmp = append(tmp, ans[i-1][j-1])
			default:
				tmp = append(tmp, ans[i-1][j-1]+ans[i-1][j])
			}
		}
		ans = append(ans, tmp)
	}
	return
}
