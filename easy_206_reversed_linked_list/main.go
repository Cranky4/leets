package main

import "fmt"

func main() {
	in := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	rev := reverseList(in)

	item := rev
	for item != nil {
		fmt.Printf("%d - ", item.Val)
		item = item.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	next := head.Next
	cur.Next = nil

	for {
		if next == nil {
			break
		}

		newNext := next.Next
		next.Next = cur

		cur = next
		next = newNext
	}

	return cur
}

type ListNode struct {
	Val  int
	Next *ListNode
}
