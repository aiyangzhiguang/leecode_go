package main

func twoSum(nums []int, target int) (ans []int) {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[target-nums[i]]
		if ok {
			return append(ans, m[target-nums[i]], i)
		}
		m[nums[i]] = i
	}
	return
}
