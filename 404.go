package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var f func(node *TreeNode, isLeft bool)
	f = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && isLeft {
			ans += node.Val
			return
		}
		f(node.Left, true)
		f(node.Right, false)

	}
	f(root, false)
	return ans
}
