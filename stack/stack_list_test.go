package stack

import (
	"fmt"
	"testing"
)

func TestStack_List(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack.Len(), stack.Empty())

	stack.Pop()
	fmt.Println(stack.Len(), stack.Empty())
}
