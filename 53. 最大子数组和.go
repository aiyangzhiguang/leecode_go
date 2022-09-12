package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	pre := 0
	ans := nums[0]
	for _, num := range nums {
		if pre > 0 {
			pre = pre + num

		} else {
			pre = num
		}
		ans = max(ans, pre)
	}
	return ans
}
