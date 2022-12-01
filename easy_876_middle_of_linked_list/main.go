package main

import "fmt"

func main() {
	// in := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	in := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}

	fmt.Printf("%v\n", middleNode(in))
}

func middleNode(head *ListNode) *ListNode {
	counter := 0

	p1 := head
	p2 := head

	for {
		if p1 == nil {
			break
		}

		p1 = p1.Next
		counter++

		if counter%2 == 0 {
			p2 = p2.Next
		}
	}

	return p2
}

type ListNode struct {
	Val  int
	Next *ListNode
}
