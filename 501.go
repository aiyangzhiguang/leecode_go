package main

import "fmt"

func findMode(root *TreeNode) []int {
	var ans []int

	var inorderTraverse func(node *TreeNode) []int
	inorderTraverse = func(node *TreeNode) []int {
		var a []int
		if node == nil {
			return a
		}
		a = append(a, inorderTraverse(node.Left)...)
		a = append(a, node.Val)
		a = append(a, inorderTraverse(node.Right)...)
		return a
	}
	m := inorderTraverse(root)
	fmt.Println(m)
	times := 0
	//ans = append(ans, root.Val)
	for i, j := 0, 1; i < j && j <= len(m); j++ {
		cur := j - i
		if j == len(m) || m[j] != m[i] {
			if cur > times {
				ans = append(ans[0:0], m[i])
			} else if cur == times {
				ans = append(ans, m[i])
			}
			i = j
			continue
		}
	}
	return ans
}
