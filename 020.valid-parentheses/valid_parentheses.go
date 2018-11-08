package main

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func isValid(s string) bool {
	stack := NewStack()
	mapping := map[rune]rune{')': '(', '}': '{', ']':'['}
	for _, char := range s {
		if _, ok := mapping[char]; ok {
			if stack.Empty() {
				return false
			}

			if topItem := stack.Pop(); topItem != mapping[char] {
				return false
			}
		} else {
			stack.Push(char)
		}
	}

	return stack.Empty()
}
