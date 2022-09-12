package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	var pre *ListNode = nil
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}
	if pre == nil {
		return slow.Next
	}
	pre.Next = slow.Next
	return head
}
