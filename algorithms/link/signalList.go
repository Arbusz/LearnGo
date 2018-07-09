package main

import (
	"fmt"
	"errors"
)

type Elem int

type LinkNode struct {
	Data Elem
	Next *LinkNode
}

func New() *LinkNode {
	return &LinkNode{0, nil}
}

func (head *LinkNode) Insert(i int, e Elem) bool {
	p := head
	j := 1
	for p != nil && j < i {
		i++
		p = p.Next
	}

	if p == nil || j > i {
		fmt.Println("越界")
		return false
	}

	s := &LinkNode{Data: e}
	s.Next = p.Next
	p.Next = s
	return true
}

func (head *LinkNode) Traverse() {
	point := head.Next
	for point != nil {
		fmt.Println(point.Data)
		point = point.Next
	}
}

func (head *LinkNode) Delate(i int) (Elem, error) {
	p := head.Next
	for j := 1; j < i; j++ {
		if p == nil {
			return -404, errors.New("越界")
		}
		p = p.Next
	}
	return p.Data, nil
}

func (head *LinkNode) Reverse() *LinkNode {
	if head.Next == nil || head == nil {
		return head
	}
	reverseHead := head
	head = head.Next
	reverseHead.Next = nil
	temp := head.Next
	for head.Next != nil {
		head.Next = reverseHead
		reverseHead = head
		head = temp
		if temp.Next != nil {
			temp = temp.Next
		}
	}
	return reverseHead
}

//递归翻转单链表

func (head *LinkNode) recursiveTraverse() *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next.recursiveTraverse() //循环到链表尾

	head.Next.Next = head //翻转链表指向
	head.Next = nil
	return newHead //新链表的头是原链表的尾
}
