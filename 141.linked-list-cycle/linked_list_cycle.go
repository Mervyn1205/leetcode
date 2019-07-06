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
 *
 * 双指针方法，一个快指针一个慢指针。如果没有环的话，最终快指针会先到达终点
 *
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}

	fmt.Println(hasCycle(head))

}
