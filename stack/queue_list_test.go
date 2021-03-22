package stack

import (
	"fmt"
	"testing"
)

func TestQueue_List(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

	for {
		e := queue.Dequeue()
		if e == nil {
			return
		}

		fmt.Println(e)
	}
}
