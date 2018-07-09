package main

import (
	"errors"
)

var (
	errNotExits     = errors.New("key is not existed")
	errTreeNil      = errors.New("tree is null")
	errTreeKeyExist = errors.New("tree index is existed")
)

type AVLNode struct {
	left   *AVLNode
	right  *AVLNode
	heigth int
	key    int
	data   Data
}

func Max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}

func getHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.heigth
}

//左单旋
//    node  BF = 2
//       \
//         pright        ----->       pright    BF = 1
//           \                        /   \
//           ppright               node  ppright
func lRotate(node *AVLNode) *AVLNode {
	pright := node.right
	node.right = pright.left
	pright.left = node

	node.heigth = Max(getHeight(node.left), getHeight(node.right)) + 1
	pright.heigth = Max(getHeight(node), getHeight(pright.right)) + 1
	return pright
}

// 右单旋
//             node  BF = -2
//              /
//         pleft         ----->        pleft   BF = 1
//            /                        /   \
//        ppleft                    ppleft   node
func rRotate(node *AVLNode) *AVLNode {
	pleft := node.left
	node.left = pleft.right
	pleft.right = node

	node.heigth = Max(getHeight(node.left), getHeight(node.left)) + 1
	pleft.heigth = Max(getHeight(node), getHeight(pleft.left)) + 1
	return pleft
}

//左右双旋
//          node                  node
//         /            左          /     右
//      node1         ---->    node2     --->         node2
//          \                   /                     /   \
//          node2            node1                 node1  node
func lrRotate(node *AVLNode) *AVLNode {
	pleft := lRotate(node.left)
	node.left = pleft
	return rRotate(node)
}

//右左双旋
//       node                  node
//          \          右         \         左
//          node1    ---->       node2     --->      node2
//          /                       \                /   \
//        node2                    node1           node  node1
func rlRotate(node *AVLNode) *AVLNode {
	pright := rRotate(node.right)
	node.right = pright
	return lRotate(node)

}

//处理节点平衡因子
func handleBF(node *AVLNode) *AVLNode {
	if getHeight(node.left)-getHeight(node.right) == 2 {
		if getHeight(node.left.left) > getHeight(node.left.right) {
			node = rRotate(node)
		} else {
			node = lrRotate(node)
		}
	} else if getHeight(node.left)-getHeight(node.right) == -2 {
		if getHeight(node.right.right) > getHeight(node.right.left) {
			node = lRotate(node)
		} else {
			node = rlRotate(node)
		}
	}

	return node
}

//插入节点，向上递归，调整平衡
func Insert(node *AVLNode, key int, data Data) (*AVLNode, error) {
	if node == nil {
		return &AVLNode{left: nil, right: nil, key: key, data: data, heigth: 1}, nil
	}
	if node.key > key {
		node.left, _ = Insert(node.left, key, data)
		node = handleBF(node)
	} else if node.key < key {
		node.right, _ = Insert(node.right, key, data)
		node = handleBF(node)
	} else {
		return nil, errTreeKeyExist
	}
	node.heigth = Max(getHeight(node.left), getHeight(node.right)) + 1
	return node, nil
}

func InorderTraverse(node *AVLNode, f func(Data)) { //函数作为参数传递,可在运行时定义匿名函数
	if node != nil {
		InorderTraverse(node.left, f)
		f(node.data)
		InorderTraverse(node.right, f)
	}
}

func searchNode(node *AVLNode, key int) (*AVLNode, bool) {
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

//删除指定的key node
//步骤，查找节点，删除节点，调整树结果
func DeleteNode(node *AVLNode, key int) (*AVLNode, error) {
	if node == nil {
		return nil, errNotExits
	}
	if node.key == key {
		if node.left == nil && node.right == nil {
			return nil, nil
		} else if node.left == nil || node.right == nil { //只空一边，哪边不空返回哪边
			if node.left != nil {
				return node.left, nil
			} else {
				return node.right, nil
			}
		} else {
			//左右都不空，寻找前驱（左支最大节点）
			//寻找左支最右
			lmr := node.left
			for {
				if lmr.right == nil {
					break
				}
				lmr = lmr.right
			}
			node.data, lmr.data = lmr.data, node.data //两个节点值进行替换
			node.key, lmr.key = lmr.key, node.key
			node.left, _ = DeleteNode(node.left, key) // 进入左子树，删掉前驱（此时已被替换）
		}
	} else if node.key > key {
		node.left, _ = DeleteNode(node.left, key)
	} else {
		node.right, _ = DeleteNode(node.right, key)
	}
	//调整高度
	node.heigth = Max(getHeight(node.left), getHeight(node.right)) + 1
	//根据平衡因子调整树
	node = handleBF(node)
	return node, nil
}
