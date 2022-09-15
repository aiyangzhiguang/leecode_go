package main

import "math"

func isValidBST(root *TreeNode) bool {
	var preNum int = math.MinInt32 - 1
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		if root.Val <= preNum {
			return false
		}
		preNum = root.Val
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return true
}
