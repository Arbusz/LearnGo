package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func breadthFirst(treeNode *Node) {
	l := list.New()
	l.PushBack(treeNode)

	for treeNode != nil {
		if l.Front() != nil {
			fmt.Printf("%v ", l.Front().Value.(*Node).data)
			treeNode = l.Front().Value.(*Node)
		} else {
			break
		}

		if treeNode.left != nil {
			l.PushBack(treeNode.left)
		}
		if treeNode.right != nil {
			l.PushBack(treeNode.right)
		}
		if l.Front() != nil {
			l.Remove(l.Front())
		}
	}
}

func Preorder(treeNode *Node) {
	if treeNode != nil {
		fmt.Printf("%v ", treeNode.data)
		if treeNode.left != nil {
			Preorder(treeNode.left)
		}
		if treeNode.right != nil {
			Preorder(treeNode.right)
		}
	} else {
		return
	}
}

func Inorder(treeNode *Node) {
	if treeNode != nil {
		if treeNode.left != nil {
			Inorder(treeNode.left)
		}
		fmt.Printf("%v ", treeNode.data)
		if treeNode.right != nil {
			Inorder(treeNode.right)
		}
	} else {
		return
	}
}

func Postorder(treeNode *Node) {
	if treeNode != nil {
		if treeNode.left != nil {
			Postorder(treeNode.left)
		}
		if treeNode.right != nil {
			Postorder(treeNode.right)
		}
		fmt.Printf("%v ", treeNode.data)
	} else {
		return
	}
}

func main() {
	node7 := Node{data: 7, left: nil, right: nil}
	node11 := Node{data: 11, left: nil, right: nil}
	node12 := Node{data: 12, left: nil, right: nil}
	node6 := Node{data: 6, left: nil, right: &node7}
	node8 := Node{data: 8, left: nil, right: nil}
	node10 := Node{data: 10, left: &node11, right: &node12}
	node4 := Node{data: 4, left: nil, right: nil}
	node5 := Node{data: 5, left: &node6, right: &node8}
	node9 := Node{data: 9, left: &node10, right: nil}
	node2 := Node{data: 2, left: &node4, right: &node5}
	node3 := Node{data: 3, left: nil, right: &node9}
	node1 := Node{data: 1, left: &node2, right: &node3}
	breadthFirst(&node1)
	fmt.Printf("\n")
	Preorder(&node1)
	fmt.Printf("\n")
	Inorder(&node1)
	fmt.Printf("\n")
	Postorder(&node1)

}
