package main

import "fmt"

// 本文件针对有头结点的链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkList(nums []int) *ListNode {
	if nums == nil || len(nums) <= 0 {
		return nil
	}
	head := ListNode{Val: nums[0]}
	p := &head
	for _, x := range nums[1:] {
		q := ListNode{Val: x}
		p.Next = &q
		p = p.Next
	}
	return &head
}

func PrintLinkedList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val)
		if p.Next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
