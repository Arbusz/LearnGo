package main

import (
	"sync"
)

type Data interface{}

type Node struct {
	key   int
	value Data
	left  *Node
	right *Node
}

type DataBST struct {
	root *Node
	lock sync.RWMutex
}

func (bst *DataBST) Insert(key int, value Data) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	newNode := &Node{key, value, nil, nil} //newNode是节点指针，不要忘记取地址

	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *DataBST) InOrderTraverse(f func(Data)) {
	bst.lock.RLock()
	defer bst.lock.Unlock()
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(node *Node, f func(Data)) { //函数作为参数传递,可在运行时定义匿名函数
	if node != nil {
		inOrderTraverse(node.left, f)
		f(node.value)
		inOrderTraverse(node.right, f)
	}
}

func (bst *DataBST) PreOrderTraverse(f func(Data)) {
	bst.lock.RLock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(node *Node, f func(Data)) {
	if node != nil {
		f(node.value)
		preOrderTraverse(node.left, f)
		preOrderTraverse(node.right, f)
	}
}

func (bst *DataBST) PostOrderTraverse(f func(Data)) {
	bst.lock.RLock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(node *Node, f func(Data)) {
	if node != nil {
		postOrderTraverse(node.left, f)
		postOrderTraverse(node.right, f)
		f(node.value)

	}
}

func (bst *DataBST) Min() *Data {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	node := bst.root
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return &node.value
}

func (bst *DataBST) Max() *Data {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	node := bst.root
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return &node.value
}

func (bst *DataBST) Search(key int) (*Node, bool) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	return search(bst.root, key)
}

func search(node *Node, key int) (*Node, bool) {
	if node == nil {
		return nil, false
	}
	if key > node.key {
		node = node.right
	}
	if key < node.key {
		node = node.left
	}
	return node, true
}

func (bst *DataBST) Remove(key int) { //删除的条目不存在不会产生错误
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = remove(node.left, key)
	}
	if key > node.key {
		node.right = remove(node.right, key)
	}
	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.right == nil {
		node = node.left
		return node
	}
	if node.left == nil {
		node = node.right
		return node
	}
	//左右都不空，寻找右支最小节点或左支最大节点
	//寻找左支最右
	lmr := node.left
	for {
		//if lmr != nil && lmr.right != nil {
		//	lmr = lmr.right
		//} else {
		//	break
		//}
		if lmr.right == nil{
			break
		}
		lmr = lmr.right
	}
	node.key, node.value = lmr.key, lmr.value //把值拷贝到节点
	node.left = remove(node.left, key)
	return node
	////寻找右支最左
	//rml := node.right
	//for {
	//	if rml != nil && rml.left != nil {
	//		rml = rml.left
	//	} else {
	//		break
	//	}
	//}
	//
	//node.key, node.value = rml.key, rml.value
	//node.right = remove(node.right, node.key)
	//return node
}
