package main

func lengthOfLongestSubstring(s string) (ans int) {
	left, right := 0, 0
	hashMap := make(map[byte]int)
	for right < len(s) {
		c := s[right]
		hashMap[c]++
		right++
		for hashMap[c] > 1 {
			d := s[left]
			left++
			hashMap[d]--
		}
		ans = max(ans, right-left)
	}
	return
}
