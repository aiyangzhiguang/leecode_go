package main

import "fmt"

//func main() {
//	var nums []interface{}
//	nums = append(nums, 1, 2, 3, nil, 4, nil, nil, 3)
//	root := CreateTree(nums)
//	inorderTraverse(root)
//	fmt.Println()
//	BFS(*root)
//}

func main() {
	var nums []interface{}
	nums = append(nums, 2, 1, 3)
	root := CreateTree(nums)
	fmt.Println(isValidBST(root))
}
