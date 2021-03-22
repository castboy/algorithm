package stack

import "container/list"

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.list.PushBack(value)
}

func (queue *Queue) Dequeue() interface{} {
	if e := queue.list.Front(); e != nil {
		queue.list.Remove(e)
		return e.Value
	}

	return nil
}

func (queue *Queue) Peek() interface{} {
	if e := queue.list.Front(); e != nil {
		return e.Value
	}

	return nil
}

func (queue *Queue) Size() int {
	return queue.list.Len()
}
