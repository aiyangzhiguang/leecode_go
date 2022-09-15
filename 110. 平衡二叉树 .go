package main

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	tmp1 := 1 + maxDepth(root.Left)
	tmp2 := 1 + maxDepth(root.Right)
	if tmp1 > tmp2 {
		return tmp1
	}
	return tmp2
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	d := math.Abs(float64(maxDepth(root.Right) - maxDepth(root.Left)))
	return d <= 1 && isBalanced(root.Right) && isBalanced(root.Left)
}
