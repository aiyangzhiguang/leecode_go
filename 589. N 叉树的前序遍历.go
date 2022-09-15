package main

type Node struct {
	Val      int
	Children []*Node
}

//func preorder(root *Node) (ans []int) {
//	if root == nil {
//		return
//	}
//	var stack []*Node
//	for root != nil || len(stack) > 0 {
//		for root != nil {
//			stack = append(stack, root)
//			if len(root.Children) > 0 {
//				root = root.Children[0]
//				root.Children = root.Children[1:]
//			} else {
//				root = nil
//			}
//		}
//		root = stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//	}
//}
