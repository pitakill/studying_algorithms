package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := NewNode(-9)
	m1 := NewNode(3)
	// n1 := NewNode(4)

	l1.AddNode(m1)
	// m1.AddNode(n1)

	l2 := NewNode(5)
	m2 := NewNode(7)
	// n2 := NewNode(4)

	l2.AddNode(m2)
	// m2.AddNode(n2)

	mergeTwoLists(l1, l2).ShowList()
}

func NewNode(i int) *ListNode {
	return &ListNode{i, nil}
}

func (n *ListNode) AddNode(m *ListNode) {
	n.Next = m
}

func (n *ListNode) ShowList() {
	if n == nil {
		return
	}

	for n.Next != nil {
		fmt.Println(n)
		n = n.Next
	}
	fmt.Println(n)
}

func AddNodeToLast(n, m *ListNode) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = m
}

func CopyNodeToLast(n, m *ListNode) {
	t := &ListNode{m.Val, nil}

	for n.Next != nil {
		n = n.Next
	}
	n.Next = t
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	t := &ListNode{}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			CopyNodeToLast(t, l1)
			l1 = l1.Next
			continue
		}

		CopyNodeToLast(t, l2)
		l2 = l2.Next
	}

	if l1 != nil {
		AddNodeToLast(t, l1)
	}
	if l2 != nil {
		AddNodeToLast(t, l2)
	}

	return t.Next
}
