package main

func inorderTraversal(root *TreeNode) (ans []int) {
	var stack []*TreeNode // 使用栈来存储已访问节点
	for root != nil || len(stack) > 0 {
		for root != nil { //如果root不为nil，一直访问到它的左节点为空
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]   // 将栈顶赋值给root
		stack = stack[:len(stack)-1] //弹出栈顶
		ans = append(ans, root.Val)
		root = root.Right
	}
	return
}
