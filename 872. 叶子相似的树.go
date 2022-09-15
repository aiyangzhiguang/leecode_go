package main

import (
	"reflect"
)

//func getLeafNums(root *TreeNode) (res []int) {
//	var ans [][]int
//	if root == nil {
//		return
//	}
//	var queue []*TreeNode
//	queue = append(queue, root)
//	for len(queue) > 0 {
//		var nodes []*TreeNode
//		var tmpAns []int
//		for len(queue) > 0 {
//			q := queue[0]
//			if q.Left != nil {
//				nodes = append(nodes, q.Left)
//			}
//			if q.Right != nil {
//				nodes = append(nodes, q.Right)
//			}
//			tmpAns = append(tmpAns, q.Val)
//			queue = queue[1:]
//		}
//		queue = append(queue, nodes...)
//		ans = append(ans, tmpAns)
//	}
//	return ans[len(ans)-1]
//}

func getLeafNums(root *TreeNode) (res []int) {
	var f func(node *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			res = append(res, root.Val)
		}
		f(root.Left)
		f(root.Right)
	}
	f(root)
	return
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	m := getLeafNums(root1)
	n := getLeafNums(root2)
	return reflect.DeepEqual(m, n)
}
