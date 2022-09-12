package main

func transformToArray(l *ListNode) []int {
	var ans []int
	for l != nil {
		ans = append(ans, l.Val)
		l = l.Next
	}
	return ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	sum := l1.Val + l2.Val + c
	if sum > 9 {
		c = 1
		sum = sum - 10
	}
	var p = &ListNode{
		Val: sum,
	}
	l1 = l1.Next
	l2 = l2.Next
	l := p
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + c
		if sum > 9 {
			c = 1
			sum = sum - 10
		} else {
			c = 0
		}
		l.Next = &ListNode{
			Val: sum,
		}
		l = l.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum = l1.Val + c
		if sum > 9 {
			c = 1
			sum = sum - 10
		} else {
			c = 0
		}
		l.Next = &ListNode{
			Val: sum,
		}
		l = l.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum = l2.Val + c
		if sum > 9 {
			c = 1
			sum = sum - 10
		} else {
			c = 0
		}
		l.Next = &ListNode{
			Val: sum,
		}
		l = l.Next
		l2 = l2.Next
	}
	if c == 1 {
		l.Next = &ListNode{
			Val: c,
		}
	}
	return p
}
