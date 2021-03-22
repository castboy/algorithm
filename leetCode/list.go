package main

import (
	"algorithm/linkedList"
)
// 删除某个链表中给定的（非末尾）节点
// ref: http://interview.wzcu.com/Leetcode/code/delete_node/
func deleteNode(node *linkedList.Node) {
	next := node.Next
	var pre *linkedList.Node

	for (next != nil) {
		node.Data = next.Data
		pre = node
		node = node.Next
		next = next.Next
	}

	pre.Next = nil
}

func main() {
	d := &linkedList.Node{Data: "d"}
	c := &linkedList.Node{Data: "c", Next: d}
	b := &linkedList.Node{Data: "b", Next: c}
	a := &linkedList.Node{Data: "a", Next: b}

	list := linkedList.List{HeadNode: a}

	list.Show()
	deleteNode(c)

	list.Show()
}
