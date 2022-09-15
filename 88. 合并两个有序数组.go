package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p, q := 0, 0
	for p < m && q < n {
		if nums1[m-1-p] >= nums2[n-1-q] {
			nums1[m+n-1-p-q] = nums1[m-1-p]
			p++
		} else {
			nums1[m+n-1-p-q] = nums2[n-1-q]
			q++
		}
	}
	for ; p < m; p++ {
		nums1[m+n-1-p-q] = nums1[m-1-p]
	}
	for ; q < n; q++ {
		nums1[m+n-1-p-q] = nums2[n-1-q]
	}
}
