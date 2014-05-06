package containers

import (
	"container/list"

	"github.com/joshuarubin/gil"
)

// Queue is a generic FIFO implementation
type Queue struct {
	list *list.List
	// tail   *LinkedListNode
	//length int
}

// Len returns the number of items in the queue
func (q *Queue) Len() int {
	if q.list == nil {
		return 0
	}

	return q.list.Len()
}

// Push an item onto the end queue
func (q *Queue) Push(item interface{}) gil.Queue {
	if q.list == nil {
		q.list = list.New()
	}

	q.list.PushBack(item)
	return q
}

// Pop an item off the front of the queue
func (q *Queue) Pop() interface{} {
	if q.list == nil || q.list.Len() == 0 {
		return nil
	}

	ret := q.list.Front()
	q.list.Remove(ret)
	return ret.Value
}

// Peek at the next item in the queue without removing it
func (q *Queue) Peek() interface{} {
	if q.list == nil || q.list.Len() == 0 {
		return nil
	}

	return q.list.Front().Value
}
