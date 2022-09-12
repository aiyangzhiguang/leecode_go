package main

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	var f func(node *TreeNode, path string)
	f = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		if len(path) == 0 {
			path = path + strconv.Itoa(node.Val)
		} else {
			path = path + "->" + strconv.Itoa(node.Val)
		}
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			return
		}
		f(node.Left, path)
		f(node.Right, path)
	}
	f(root, "")
	return paths
}
