package main

//You are given two non-empty linked lists representing two non-negative integers.
//The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carryFlag := 0

	var headNode *ListNode
	var tailNode *ListNode

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carryFlag

		if sum >= 10 {
			carryFlag = 1
			sum = sum % 10
		} else {
			carryFlag = 0
		}

		node := &ListNode{
			sum,
			nil,
		}

		if headNode == nil {
			headNode = node
			tailNode = node
		} else {
			tailNode.Next = node
			tailNode = node
		}
	}

	if carryFlag == 1 {
		tailNode.Next = &ListNode{
			carryFlag,
			nil,
		}
	}

	return headNode
}
