package tree

import "fmt"

type Node struct {
	Val string
	Left, Right *Node
}

// 深度优先遍历
// 递归实现

func PreOrder(root *Node) {
	fmt.Println(root.Val)

	if root.Left != nil {
		PreOrder(root.Left)
	}

	if root.Right != nil {
		PreOrder(root.Right)
	}
}

func InOrder(root *Node) {
	if root.Left != nil {
		InOrder(root.Left)
	}

	fmt.Println(root.Val)

	if root.Right != nil {
		InOrder(root.Right)
	}
}

func PostOrder(root *Node) {
	if root.Left != nil {
		PostOrder(root.Left)
	}

	if root.Right != nil {
		PostOrder(root.Right)
	}

	fmt.Println(root.Val)
}

// 深度优先遍历
// 堆栈实现


