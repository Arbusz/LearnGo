package main

import (
	"errors"
)

type Stack []interface{} //空接口类型的切片

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) IsEmpty() bool {
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("The stack is empty")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack //值拷贝
	if stack.IsEmpty() {
		return nil, errors.New("The stack is empty")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	
	return value, nil
}

func (stack *Stack) Push(data interface{}) {
	*stack = append(*stack, data)
}
