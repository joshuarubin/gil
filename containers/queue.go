package containers

import (
	"github.com/joshuarubin/gil"
)

// Queue is a generic FIFO implementation
type Queue struct {
	list   *LinkedList
	tail   *LinkedListNode
	length int
}

// Len returns the number of items in the queue
func (q *Queue) Len() int {
	return q.length
}

// Push an item onto the end queue
func (q *Queue) Push(item interface{}) gil.Queue {
	q.length++

	if q.list == nil {
		q.list = &LinkedList{}
		q.list.Push(item)
		q.tail = q.list.Head
		return q
	}

	q.tail = q.tail.InsertNext(item)
	return q
}

// Pop an item off the front of the queue
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}

	q.length--
	return q.list.Pop()
}

// Peek at the next item in the queue without removing it
func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}

	return q.list.Head.Value
}
