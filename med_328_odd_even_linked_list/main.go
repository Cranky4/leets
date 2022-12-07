package main

import "fmt"

func main() {
	in := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	head := oddEvenList(in)

	for {
		fmt.Printf("%v ", head.Val)

		if head.Next == nil {
			break
		}

		head = head.Next
	}
}

func oddEvenList(head *ListNode) *ListNode {
	n := head

	var odds *ListNode
	var evens *ListNode
	var firstOdd *ListNode
	var firstEven *ListNode

	i := 1
	for {
		if n == nil {
			break
		}

		if i%2 == 1 { // odd
			if odds == nil {
				odds = n
				firstOdd = n
			} else {
				odds.Next = n
				odds = odds.Next
			}
		} else {
			if evens == nil {
				evens = n
				firstEven = n
			} else {
				evens.Next = n
				evens = evens.Next
			}
		}

		n = n.Next

		i++
	}

	evens.Next = nil
	odds.Next = firstEven

	return firstOdd
}

type ListNode struct {
	Val  int
	Next *ListNode
}
