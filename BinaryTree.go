package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := TreeNode{Val: nums[0].(int)}
	var queue []*TreeNode
	queue = append(queue, &root)
	for i := 1; len(queue) > 0 && i < len(nums); {
		cur := queue[0]
		if nums[i] != nil {
			cur.Left = &TreeNode{
				Val: nums[i].(int),
			}
			queue = append(queue, cur.Left)
		}
		i++
		if i < len(nums) {
			if nums[i] != nil {
				cur.Right = &TreeNode{
					Val: nums[i].(int),
				}
				queue = append(queue, cur.Right)
			}
			i++
		}
		queue = queue[1:]
	}
	return &root
}

func inorderTraverse(root *TreeNode) {

	if root == nil {
		return
	}
	inorderTraverse(root.Left)
	fmt.Print(root.Val)
	fmt.Print("\t")
	inorderTraverse(root.Right)
}

func BFS(root TreeNode) {
	if &root == nil {
		return
	}
	var queue []TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		if cur.Left != nil {
			queue = append(queue, *cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, *cur.Right)
		}
		queue = queue[1:]
		fmt.Print(strconv.Itoa(cur.Val) + "\t")
	}
}
