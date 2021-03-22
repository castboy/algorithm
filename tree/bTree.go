package tree

import (
	"algorithm/stack"
	"fmt"
)

type Node struct {
	Val string
	Left, Right *Node
}

// 深度优先遍历
// 递归实现

func PreOrder(root *Node) {
	//fmt.Println(root.Val)
	//
	//if root.Left != nil {
	//	PreOrder(root.Left)
	//}
	//
	//if root.Right != nil {
	//	PreOrder(root.Right)
	//}

	if root == nil {
		return
	}

	fmt.Print(root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func InOrder(root *Node) {
	//if root.Left != nil {
	//	InOrder(root.Left)
	//}
	//
	//fmt.Println(root.Val)
	//
	//if root.Right != nil {
	//	InOrder(root.Right)
	//}

	if root == nil {
		return
	}

	InOrder(root.Left)
	fmt.Print(root.Val)
	InOrder(root.Right)
}

func PostOrder(root *Node) {
	//if root.Left != nil {
	//	PostOrder(root.Left)
	//}
	//
	//if root.Right != nil {
	//	PostOrder(root.Right)
	//}
	//
	//fmt.Println(root.Val)

	if root == nil {
		return
	}

	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Print(root.Val)
}

func leftPush(stack *stack.Stack, root *Node) {
	for {
		if root != nil {
			stack.Push(*root)
			fmt.Print(root.Val)
			root = root.Left
		} else {
			break
		}
	}
}

func PreOrderStack(root *Node) {
	bStack := stack.NewStack()

	leftPush(bStack, root)

	for {
		if e := bStack.Pop(); e != nil {
			leftPush(bStack, e.(Node).Right)
		} else {
			break
		}
	}
}

func rightPush(stack *stack.Stack, root *Node) {
	for {
		if root != nil {
			stack.Push(*root)
			root = root.Right
		} else {
			break
		}
	}
}

func InOrderStack(root *Node) {
	aStack := stack.NewStack()
	bStack := stack.NewStack()

	rightPush(bStack, root)

	for {
		if e := bStack.Pop(); e != nil {
			aStack.Push(e)
			rightPush(bStack, e.(Node).Left)
		} else {
			break
		}
	}

	for {
		if e := aStack.Pop(); e != nil {
			fmt.Print(e.(Node).Val)
		} else {
			break
		}
	}
}
// ref: https://blog.csdn.net/sjf0115/article/details/8645991
func PreOrder2(root *Node) {
	aStack := stack.NewStack()
	node := root

	for (node != nil || !aStack.Empty()) { // node != nil -> just judge root node
		if (node != nil) {
			aStack.Push(node)
			fmt.Println(node.Val)
			node = node.Left
		} else {
			newRoot := aStack.Pop().(*Node) // 1：left节点为nil，则弹出；2：弹出的right节点为nil，接着弹出；
			node = newRoot.Right // 每次弹出向right节点走
		}
	}
}

func InOrder2(root *Node) {
	aStack := stack.NewStack()
	node := root

	for (node != nil || !aStack.Empty()) { // node != nil -> just judge root node
		if (node != nil) {
			aStack.Push(node)
			node = node.Left
		} else {
			newRoot := aStack.Pop().(*Node)
			fmt.Println(newRoot.Val)
			node = newRoot.Right
		}
	}
}

// ref: https://zoyi14.smartapps.cn/pages/note/index?slug=456af5480cee&origin=share&hostname=baiduboxapp&_swebfr=1
func PostOrder2(root *Node) {
	aStack := stack.NewStack()
	node := root
	lastVisit := root

	for (node != nil || !aStack.Empty()) { // node != nil -> just judge root node
		for (node != nil) { // step 1: just push left child node.
			aStack.Push(node)
			node = node.Left
		}

		node = aStack.Peek().(*Node) //step 2: can not Pop if node.Right is not nil or not visited.

		if (node.Right == nil || node.Right == lastVisit) {
			fmt.Println(node.Val)
			aStack.Pop()
			lastVisit = node
			node = nil // just escape step 1 to step 2 directly.
		} else {
			node = node.Right // as a new root node, exec step 1
		}
	}
}

