package main

import (
	"testing"
)

func TestmergeTwoListsRecursive(t *testing.T) {
	n1 := newNode(1)
	n1.Next = newNode(2)
	n1.Next.Next = newNode(4)

	n2 := newNode(1)
	n2.Next = newNode(3)
	n2.Next.Next = newNode(4)

	list := mergeTwoListsRecursive(n1, n2)

	if list.Val != 1 {
		t.Error("Val should be 1")
	}

	if list.Next.Val != 1 {
		t.Error("Val should be 1")
	}

	if list.Next.Next.Val != 2 {
		t.Error("Val should be 2")
	}

	if list.Next.Next.Next.Val != 3 {
		t.Error("Val should be 3")
	}

	if list.Next.Next.Next.Next.Val != 4 {
		t.Error("Val should be 4")
	}

	if list.Next.Next.Next.Next.Next.Val != 4 {
		t.Error("Val should be 4")
	}

}


func TestmergeTwoListsIterative(t *testing.T) {
	n1 := newNode(0)
	n1.Next = newNode(2)
	n1.Next.Next = newNode(4)

	n2 := newNode(1)
	n2.Next = newNode(3)
	n2.Next.Next = newNode(5)

	list := mergeTwoListsIterative(n1, n2)

	if list.Val != 0 {
		t.Error("Val should be 0")
	}

	if list.Next.Val != 1 {
		t.Error("Val should be 1")
	}

	if list.Next.Next.Val != 2 {
		t.Error("Val should be 2")
	}

	if list.Next.Next.Next.Val != 3 {
		t.Error("Val should be 3")
	}

	if list.Next.Next.Next.Next.Val != 4 {
		t.Error("Val should be 4")
	}

	if list.Next.Next.Next.Next.Next.Val != 5 {
		t.Error("Val should be 5")
	}
}
