package main

func canJump(nums []int) bool {
	max := 0
	for i, jump := range nums {
		if max >= i && i+jump > max {
			max = i + jump
		}
	}
	return max >= len(nums)-1
}
