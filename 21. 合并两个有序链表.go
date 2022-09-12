package main

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	p := list1
	q := list2
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	var preHead = &ListNode{
		Val: -1,
	}
	prev := preHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			prev.Next = p
			p = p.Next
		} else {
			prev.Next = q
			q = q.Next
		}
		prev = prev.Next
	}
	if p != nil {
		prev.Next = p
	} else {
		prev.Next = q
	}
	return preHead.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
