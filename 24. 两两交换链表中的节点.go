package main

// 迭代方法
//func swapPairs(head *ListNode) *ListNode {
//	pre := ListNode{
//		Val:  0,
//		Next: head,
//	}
//	temp := &pre
//	for temp.Next != nil && temp.Next.Next != nil {
//		first := temp.Next
//		second := first.Next
//		temp.Next = second
//		first.Next = second.Next
//		second.Next = first
//		temp = first
//	}
//	return pre.Next
//}

// 递归方法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next
	head.Next = swapPairs(head.Next.Next)
	tmp.Next = head
	return tmp
}
