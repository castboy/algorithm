package tree

import "testing"

func newBTree() *Node {
	root := &Node{Val: "A"}

	root.Left = &Node{Val: "B"}
	root.Left.Left = &Node{Val: "D"}
	root.Left.Right = &Node{Val: "E"}

	root.Right = &Node{Val: "C"}
	root.Right.Left = &Node{Val: "F"}
	root.Right.Right = &Node{Val: "G"}

	return root
}

func TestPreOrder(t *testing.T) {
	root := newBTree()
	PreOrder(root)
}

func TestInOrder(t *testing.T) {
	root := newBTree()
	InOrder(root)
}

func TestPostOrder(t *testing.T) {
	root := newBTree()
	PostOrder(root)
}

func TestInOrderStack(t *testing.T) {
	root := newBTree()
	InOrderStack(root)
}

func TestPreOrderStack(t *testing.T) {
	root := newBTree()
	PreOrderStack(root)
}