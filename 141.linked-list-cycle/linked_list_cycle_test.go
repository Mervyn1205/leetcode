package main

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}

	if hasCycle(head) {
		t.Error("[1, 2] has not cycle")
	}

	node1 := ListNode{3, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{0, nil}
	node4 := ListNode{-4, nil}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node2

	if !hasCycle(&node1) {
		t.Error("[3, 2, 0, -4] has cycle")
	}

}
