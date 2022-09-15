package main

func permute(nums []int) (ans [][]int) {
	var traceback func()
	visited := make([]bool, len(nums))
	var path []int
	traceback = func() {
		if len(path) == len(nums) {
			var temp = make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i, x := range nums {
			if visited[i] {
				continue
			}
			path = append(path, x)
			visited[i] = true
			traceback()
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	traceback()
	return
}
