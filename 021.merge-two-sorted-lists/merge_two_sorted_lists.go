package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursive(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsRecursive(l1, l2.Next)
		return l2
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoListsIterative(l1 *ListNode, l2 *ListNode) *ListNode {
	list := new(ListNode)
	head := list
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}

		head = head.Next
	}

	if l1 == nil {
		head.Next = l2
	} else {
		head.Next = l1
	}

	return list.Next
}

func newNode(val int) *ListNode {
	node := new(ListNode)
	node.Val = val

	return node
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
func main() {
	n1 := newNode(1)
	n1.Next = newNode(2)
	n1.Next.Next = newNode(4)

	n2 := newNode(1)
	n2.Next = newNode(3)
	n2.Next.Next = newNode(4)

	printList(mergeTwoListsRecursive(n1, n2))

	n3 := newNode(0)
	n3.Next = newNode(2)
	n3.Next.Next = newNode(4)

	n4 := newNode(1)
	n4.Next = newNode(3)
	n4.Next.Next = newNode(5)

	printList(mergeTwoListsIterative(n3, n4))


}
