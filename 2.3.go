package main

import (
	"fmt"
	"os"
)

const (
	ERROR = -1000000001
)

type Element int64

type LinkNode struct {
	Data Element
	Nest *LinkNode
}

type LinkNoder interface {
	Add(head *LinkNode, new *LinkNode)              //增加
	Delete(head *LinkNode, index int)               //删除
	Insert(head *LinkNode, index int, data Element) //插入
	GetLength(head *LinkNode) int                   //获取长度
	GetData(head *LinkNode, index int) Element      //获取指定index位置的元素
	ReverseList(head *LinkNode)                     //翻转
	RecursiveTraverse(head *LinkNode)               //递归翻转
	Traverse(head *LinkNode)
}

var newList *LinkNode
var endNode *LinkNode

func Add(head *LinkNode, data Element) {
	point := head
	for point.Nest != nil {
		point = point.Nest
	}
	var node LinkNode
	point.Nest = &node
	node.Data = data
	head.Data = Element(GetLength(head))

	if GetLength(head) > 1 {
		Traverse(head)
	}
}

func Delete(head *LinkNode, index int) Element {
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return ERROR
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest
		}
		point.Nest = point.Nest.Nest
		data := point.Nest.Data
		return data
	}
}

func Insert(head *LinkNode, index int, data Element) {
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		//return ERROR
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest
		}
		var node LinkNode
		node.Data = data
		node.Nest = point.Nest
		point.Nest = &node
	}
}

func GetLength(head *LinkNode) int {
	point := head
	var length int
	for point.Nest != nil {
		length++
		point = point.Nest
	}
	return length
}

func GetData(head *LinkNode, index int) Element {
    point := head
    if index < 0 || index > GetLength(head) {
        fmt.Println("please check index")
        return ERROR
    } else {
        for i := 0; i < index; i++ {
            point = point.Nest
        }
        return point.Data
    }
}

func ReverseList(head *LinkNode) *LinkNode {
	RecursiveTraverse(head)

	return newList
}

func RecursiveTraverse(head *LinkNode) {
	if head == nil {
		return
	}
	if head.Nest == nil {
		endNode = head
		newList = endNode
		return
	}

	RecursiveTraverse(head.Nest)

	endNode.Nest = head
	endNode = head
	endNode.Nest = nil
}


func Traverse(head *LinkNode) {
    point := head.Nest
    for point.Nest != nil {
        fmt.Println(point.Data)
        point = point.Nest
    }
    fmt.Println("Traverse OK!")
}




func main() {
	var head LinkNode = LinkNode{Data: 0, Nest: nil}
	head.Data = 0
	var nodeArray []Element
	for i := 0; i < 10; i++ {
		nodeArray = append(nodeArray, Element(i+1+i*100))
		Add(&head, nodeArray[i])
	}
	fmt.Println("data:", GetData(&head, 1))

	point := ReverseList(&head)
	//fmt.Println("data:", GetData(point, 1))
	for j:= 0; j <10; j++{
		fmt.Println(GetData(point, j))
	}

	fmt.Println("length:", GetLength(point))

	os.Exit(0)
}