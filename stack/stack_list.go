package stack

import "container/list"

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (stack *Stack) Push(value interface{}) {
	stack.l.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	if e := stack.l.Back(); e != nil {
		stack.l.Remove(e)
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.l.Len()
}

func (stack *Stack) Empty() bool {
	return stack.l.Len() == 0
}