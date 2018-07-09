package main

func (bst *DataBST) PreOrder(f func(Data)) {
	bst.lock.RLock()
	defer bst.lock.Unlock()
	preOrder(bst.root, f)
}

func preOrder(node *Node, f func(Data)) { //函数作为参数传递,可在运行时定义匿名函数
	if node != nil {
		stack := new(Stack)
		stack.Push(node)
		for !stack.IsEmpty() {
			p := stack.Pop().(*Node) //这他妈是个啥，为啥这样写就好使了
			f(p.value)
			if p.left != nil {
				stack.Push(p.right)
			}
			if p.right != nil {
				stack.Push(p.left)
			}

		}
	}
}
func (bst *DataBST) InOrder(f func(Data)) {
	bst.lock.RLock()
	defer bst.lock.Unlock()
	inOrder(bst.root, f)
}

func inOrder(node *Node, f func(Data)) {
	if node != nil {
		stack := new(Stack)
		p := node
		for p != nil || !stack.IsEmpty() {
			if p != nil {
				stack.Push(p) //进入左子树一直访问到最左节点
				p = p.left
			} else {
				p = stack.Pop().(*Node)
				f(p.value)  //访问数据
				p = p.right //进入右子树
			}
		}
	}
}

func (bst *DataBST) PostOrder(f func(Data)) {
	bst.lock.RLock()
	defer bst.lock.Unlock()
	postOrder(bst.root, f)
}

//有点迷
//func postOrder(node *Node, f func(Data)) {
//	S := new([100]*Node)
//	Tag := new([100]int)
//	for p, top := node, 0; p != nil || top != 0; {
//		for p != nil {
//			S[top] = p
//			Tag[top] = 0
//			top ++
//			p = p.left
//
//		}
//		for top != 0 && Tag[top] == 1 {
//			p = S[top]
//			top --
//			f(p.value)
//		}
//		if top != 0 {
//			Tag [top] = 1
//			p = S[top].right
//		}
//	}
//
//}
//更迷了
func postOrder(node *Node, f func(Data)) {
	if node != nil {
		stack := new(Stack)

		var preVisited *Node
		for p := node; p != nil || !stack.IsEmpty(); {
			for p != nil {
				stack.Push(p)
				p = p.left
			}

			top := stack.Pop().(*Node)
			if (top.left == nil && top.right == nil) || (top.right == nil && preVisited == top.left) || preVisited == top.right {
				f(top.value)
				preVisited = top
			} else {
				p = top.right
			}
		}

	}
}
