package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, x := range nums1 {
		m[x]++
	}
	ans := nums1[0:0]
	for _, x := range nums2 {
		if m[x] > 0 {
			ans = append(ans, x)
			m[x]--
		}
	}
	return ans
}
