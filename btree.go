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

type seqStack struct {
	data [100]*Node
	tag  [100]int // 后续遍历准备
	top  int      // 数组下标
}

func treeMake(data []interface{}) *Node {
	if len(data) == 0 {
		return nil
	}
	root := &Node{}
	switch t := data[0].(type) {
	case int:
		root.data = t
	case nil:
		return nil
	default:
		panic("Unknown element type")
	}

	queue := make([]*Node, 1)
	queue[0] = root

	data = data[1:]
	for len(data) > 0 && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 左侧节点
		node.left = newNodeFromData(data[0])
		if node.left != nil {
			queue = append(queue, node.left)
		}
		data = data[1:]

		// 右侧节点
		if len(data) > 0 {
			node.right = newNodeFromData(data[0])
			if node.right != nil {
				queue = append(queue, node.right)
			}
			data = data[1:]
		}
	}
	return root
}

func newNodeFromData(val interface{}) *Node {
	switch t := val.(type) {
	case int:
		return &Node{data: t}
	case nil:
		return nil
	default:
		panic("Unknown element type")
	}
}

//func read() []Node {
//	var N int
//	fmt.Scanf("%d", &N)
//	var nodes []Node = make([]Node, N)
//	for i := 0; i < N; i++ {
//		var da, indexLeft, indexRight int
//		fmt.Scanf("%d %d %d",&da,&indexLeft,&indexRight)
//		nodes[i].data=da
//		if indexLeft >= 0 {
//			nodes[i].left = &nodes[indexLeft]
//		}
//		if indexRight >= 0{
//			nodes[i].right = &nodes[indexRight]
//		}
//	}
//	return nodes
//}

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
//
//func Preorder(node *Node) (result []int) {
//	var s seqStack
//	s.top = -1 // 空
//	if node == nil {
//		panic("no data here")
//	}else {
//		for node != nil || s.top != -1 {
//			for node != nil {
//				result = append(result, node.data)
//				s.top++
//				s.data[s.top] = node
//				node = node.left
//			}
//			s.top--
//			node = s.data[s.top + 1]
//			node = node.right
//		}
//	}
//	fmt.Println(result)
//	return
//}
//
//func Inorder(node *Node) (result []int ) {
//	var s seqStack
//	s.top = -1
//	if node == nil {
//		panic("no data here")
//	}else {
//		for node != nil || s.top != -1 {
//			for node != nil {
//				s.top++
//				s.data[s.top] = node
//				node = node.left
//			}
//			s.top--
//			node = s.data[s.top + 1]
//			result = append(result, node.data)
//			node = node.right
//		}
//	}
//	fmt.Println(result)
//	return
//}
//
//
//func Postorder(node *Node) (result []int)  {
//	var s seqStack
//	s.top = -1
//
//	if node == nil {
//		panic("no data here")
//	}else {
//		for node != nil || s.top != -1 {
//			for node != nil {
//				s.top++
//				s.data[s.top] = node
//				s.tag[s.top] = 0
//				node = node.left
//			}
//
//			if s.tag[s.top] == 0 {
//				node = s.data[s.top]
//				s.tag[s.top] = 1
//				node = node.right
//			}else {
//				for s.tag[s.top] == 1 {
//					s.top--
//					node = s.data[s.top + 1]
//					result = append(result, node.data)
//					if s.top < 0 {
//						break
//					}
//				}
//				node = nil
//			}
//		}
//	}
//	fmt.Println(result)
//	return
//}

func main() {
	treeRoot := treeMake([] interface{}{1, 2, 3, nil, nil, 4, 5})
	breadthFirst(treeRoot)
	fmt.Printf("\n")
	Preorder(treeRoot)
	fmt.Printf("\n")
	Inorder(treeRoot)
	fmt.Printf("\n")
	Postorder(treeRoot)

}

