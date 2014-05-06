package queue

import (
	"container/list"

	"github.com/joshuarubin/gil"
)

// Stack is a generic LIFO implementation
type Stack struct {
	list *list.List
}

// Len returns the number of items in the stack
func (s *Stack) Len() int {
	if s.list == nil {
		return 0
	}

	return s.list.Len()
}

// Push an item at the front of the linked list
func (s *Stack) Push(item gil.Interface) gil.Queue {
	if s.list == nil {
		s.list = list.New()
	}

	s.list.PushBack(item)
	return s
}

// Pop an item off the front of the linked list
func (s *Stack) Pop() gil.Interface {
	if s.list == nil || s.list.Len() == 0 {
		return nil
	}

	ret := s.list.Back()
	s.list.Remove(ret)
	return ret.Value.(gil.Interface)
}

// Peek at the item at the front of the stack without removing it
func (s *Stack) Peek() gil.Interface {
	if s.list == nil || s.list.Len() == 0 {
		return nil
	}

	return s.list.Back().Value.(gil.Interface)
}
