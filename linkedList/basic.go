package linkedList

import "fmt"

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	HeadNode *Node
}

func (l *List) IsEmpty() bool {
	if l.HeadNode == nil {
		return true
	}

	return false
}

func (l *List) Length() int {
	var count int
	cur := l.HeadNode

	for cur != nil {
		count++
		cur = cur.Next
	}

	return count
}

func (l *List) Add(data Object) {
	newHeadNode := &Node{
		Data: data,
	}

	newHeadNode.Next = l.HeadNode
	l.HeadNode = newHeadNode
}

func (l *List) Append(data Object) {
	node := &Node{
		Data: data,
	}

	if l.HeadNode == nil {
		l.HeadNode = node
		return
	}

	cur := l.HeadNode

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = node
}

func (l *List) Insert(index int, data Object) {
	if index <= 0 {
		l.Add(data)
		return
	}

	cur := l.HeadNode
	count := 0

	for cur != nil {
		count++

		if count == index {
			//next := cur.Next
			//cur.Next = &Node{Data:data}
			//cur.Next.Next = next
			node := &Node{Data:data}
			node.Next = cur.Next
			cur.Next = node

			return
		}

		cur = cur.Next
	}

	l.Append(data)
}

func (l *List) Remove(data Object) {
	if data == l.HeadNode.Data {
		l.HeadNode = l.HeadNode.Next
		return
	}

	cur := l.HeadNode

	for cur.Next != nil {
		if cur.Next.Data == data {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
	}
}

func (l *List) RemoveAtIndex(index int) {
	if index <= 0 {
		l.HeadNode = l.HeadNode.Next
		return
	}

	cur := l.HeadNode
	count := 0

	for cur != nil {
		count++

		if index == count {
			if cur.Next != nil {
				cur.Next = cur.Next.Next
			}

			return
		}

		cur = cur.Next
	}
}

func (l *List) Contain(data Object) bool {
	cur := l.HeadNode

	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}

	return false
}

func (l *List) Reverse() {
	var newHeadNode *Node
	var node *Node

	for l.HeadNode != nil {
		node = l.HeadNode
		l.HeadNode = l.HeadNode.Next

		node.Next = newHeadNode
		newHeadNode = node
	}

	l.HeadNode = newHeadNode
}

func (l *List) Show() {
	if l.HeadNode == nil {
		fmt.Println("nil linkedList")
	}

	var cur = l.HeadNode

	for cur != nil {
		fmt.Println(cur.Data.(string))
		cur = cur.Next
	}

	fmt.Println("-------------------")
}