package main

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var nextQueue []*TreeNode
		var tmp []int
		for len(queue) > 0 {
			cur := queue[0]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				nextQueue = append(nextQueue, cur.Left)
			}
			if cur.Right != nil {
				nextQueue = append(nextQueue, cur.Right)
			}
			queue = queue[1:]
		}
		ans = append(ans, tmp)
		queue = append(queue, nextQueue...)
	}
	return ans
}
