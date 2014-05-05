package containers

import (
	"fmt"
	"strings"
)

// LinkedListNode is a node in a linked list
type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

// InsertNext inserts item after the current node.
// Returns a pointer to the node that was just inserted.
func (node *LinkedListNode) InsertNext(item interface{}) *LinkedListNode {
	newNode := &LinkedListNode{
		Value: item,
		Next:  node.Next,
	}
	node.Next = newNode
	return newNode
}

// RemoveNext removes the node after the current node.
// Returns a pointer to the current node.
func (node *LinkedListNode) RemoveNext() *LinkedListNode {
	if node.Next != nil {
		node.Next = node.Next.Next
	}
	return node
}

// LinkedList is an object that contains the Head pointer of a linked list
// as well as methods that apply to the whole list
type LinkedList struct {
	Head *LinkedListNode
}

// Push inserts item at the beginning of the linked list
func (l *LinkedList) Push(item interface{}) *LinkedList {
	l.Head = &LinkedListNode{
		Value: item,
		Next:  l.Head,
	}
	return l
}

// Pop removes and returns the item at the beginning of the linked list
func (l *LinkedList) Pop() interface{} {
	node := l.Head
	if node == nil {
		return nil
	}

	l.Head = node.Next
	return node.Value
}

// Reverse reverses the linked list beginning at Head
func (l *LinkedList) Reverse() {
	var prev, next *LinkedListNode
	cur := l.Head

	for cur != nil {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	l.Head = prev
}

func (l *LinkedList) String() string {
	ret := ""

	l.Walk(func(node *LinkedListNode, _ int) bool {
		ret = fmt.Sprintf("%s %v", ret, node.Value)
		return true
	})

	return strings.TrimSpace(ret)
}

// Walk executes f for each node of the linked list beginning at Head.
// f is passed a pointer to a node and the index of that node within the list.
// The Walk stops if f returns false or when the end of the list is reached.
// Walk returns the last node which will be nil if the end of the list is reached or
// the node that f returned false for.
// It will also return the index of that node.
func (l *LinkedList) Walk(f func(*LinkedListNode, int) bool) (*LinkedListNode, int) {
	node := l.Head
	i := 0
	for node != nil {
		if !f(node, i) {
			break
		}
		node = node.Next
		i++
	}
	return node, i
}
