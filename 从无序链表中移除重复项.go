package main

func RemoveDuplicatedNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	m := make(map[int]int)
	pre := head.Next
	m[(*pre).Val] = 1
	p := head.Next.Next
	for p != nil {
		_, ok := m[p.Val]
		if ok {
			pre.Next = p.Next
		} else {
			m[(*p).Val] = 1
		}
		p = p.Next
		pre = pre.Next
	}
	return head
}
