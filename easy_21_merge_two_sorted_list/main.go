package main

import "fmt"

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	fmt.Printf("%v\n", mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	start, list1, list2 := getNext(list1, list2)

	if start == nil {
		return nil
	}

	cur := start
	next, list1, list2 := getNext(list1, list2)

	for next != nil {
		cur.Next = next
		cur = cur.Next
		next, list1, list2 = getNext(list1, list2)
	}

	return start
}

func getNext(node1 *ListNode, node2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	if node1 == nil && node2 == nil {
		return nil, nil, nil
	}

	if node1 == nil && node2 != nil {
		return node2, node1, node2.Next
	}

	if node2 == nil && node1 != nil {
		return node1, node1.Next, node2
	}

	if node1.Val < node2.Val {
		return node1, node1.Next, node2
	}

	return node2, node1, node2.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
