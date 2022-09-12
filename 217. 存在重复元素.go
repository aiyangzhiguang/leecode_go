package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]byte)
	for _, x := range nums {
		_, ok := m[x]
		if ok {
			return true
		}
		m[x] = 1
	}
	return false
}
