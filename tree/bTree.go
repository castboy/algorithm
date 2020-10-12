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


