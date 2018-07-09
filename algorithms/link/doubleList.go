package main

type DNode struct {
	data Elem
	prev *DNode
	next *DNode
}

type DList struct {
	size int
	head *DNode
	tail *DNode
}

func (dlist *DList) Init() {
	_dList := dlist
	_dList.size = dlist.size
	_dList.head = nil
	_dList.tail = nil
}

func (dList *DList) Append(data Elem) {
	newNode := new(DNode)
	newNode.data = data

	if dList.size == 0 {
		dList.head = newNode
		dList.tail = newNode
		newNode.prev = nil
		newNode.next = nil
	} else {
		newNode.prev = dList.tail
		newNode.next = nil
		dList.tail = newNode
		dList.tail.next = newNode
	}
	dList.size++
}

func (dList *DList) InsertNext(tag *DNode, data Elem) bool {
	if tag == nil {
		return false
	}

	if dList.isTail(tag) {
		dList.Append(data)
	} else {
		newNode := new(DNode)
		newNode.data = data
		newNode.prev = tag
		newNode.next = tag.next

		tag.next = newNode
		newNode.next.prev = newNode
		dList.size++
	}
	return true
}

func (dList *DList) InsertPrev(tag *DNode, data Elem) bool {
	if tag == nil {
		return false
	}

	if dList.head == tag {
		newNode := new(DNode)
		newNode.data = data
		newNode.next = dList.head
		newNode.prev = nil

		dList.head.prev = newNode
		dList.head = newNode
		dList.size++
		return true
	} else {
		prev := tag.prev //前插可以理解为前一个节点的后插
		return dList.InsertNext(prev, data)
	}

}

func (dList *DList) isHead(dNode *DNode) bool {
	if dNode == dList.head {
		return true
	} else {
		return false
	}
}

func (dList *DList) isTail(dNode *DNode) bool {
	if dNode == dList.tail {
		return true
	} else {
		return false
	}
}

func (dList *DList) Remove(tag *DNode) Elem {
	if tag == nil {
		return -404
	}

	prev := tag.prev
	next := tag.next

	if dList.isHead(tag) {
		dList.head = next
	} else {
		prev.next = next
	}

	if dList.isTail(tag) {
		dList.tail = prev
	} else {
		next.prev = prev
	}

	dList.size--
	return tag.data
}

func (dList *DList) Search(data Elem) *DNode {
	p := dList.head
	for p != nil {
		if p.data == data {
			return p
		}
	}

	return nil
}
