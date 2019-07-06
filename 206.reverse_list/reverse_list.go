package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * 在遍历列表时，将当前节点的 next 指针改为指向前一个元素。由于节点没有引用其上一个节点，因此必须事先存储其前一个元素。在更改引用之前，还需要另一个指针来存储下一个节点。不要忘记在最后返回新的头引用！
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prevNode *ListNode
	currNode := head


	for currNode != nil {
		nextTempNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextTempNode
	}

	return prevNode
}

func newNode(val int) *ListNode {
	node := new(ListNode)
	node.Val = val

	return node
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d  ", list.Val)
		list = list.Next
	}

	fmt.Printf("\n")
}

func main()  {
	n1 := newNode(1)
	n1.Next = newNode(2)
	n1.Next.Next = newNode(3)
	n1.Next.Next.Next = newNode(4)
	n1.Next.Next.Next.Next = newNode(5)

	printList(n1)
	printList(reverseList(n1))
}
