package main

func getAllNum(num int) []int {
	var nums []int
	for num != 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	return nums
}

//func maximumSwap(num int) int {
//
//}
