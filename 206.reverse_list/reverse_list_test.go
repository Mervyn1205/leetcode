package main

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	n1 := newNode(1)
	n1.Next = newNode(2)
	n1.Next.Next = newNode(3)


	list := reverseList(n1)

	if list.Val != 3 {
		t.Error("Val should be 3")
	}

	if list.Next.Val != 2 {
		t.Error("Val should be 2")
	}

	if list.Next.Next.Val != 1 {
		t.Error("Val should be 1")
	}

}