package main

import "math"

func isValidBST(root *TreeNode) bool {
	var preNum int = math.MinInt32
	var stack []*TreeNode
	stack = append(stack, root)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		println(preNum)
		print("root.val:")
		println(root.Val)
		if root.Val < preNum {
			return false
		}
		preNum = root.Val
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return true
}
