package containers

import (
	"github.com/joshuarubin/gil"
)

// Stack is a generic LIFO implementation
type Stack struct {
	list   *LinkedList
	length int
}

// Len returns the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

// Push an item at the front of the linked list
func (s *Stack) Push(item interface{}) gil.Queue {
	s.length++

	if s.list == nil {
		s.list = &LinkedList{}
	}

	s.list.Push(item)
	return s
}

// Pop an item off the front of the linked list
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	s.length--
	return s.list.Pop()
}

// Peek at the item at the front of the stack without removing it
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}

	return s.list.Head.Value
}
