package containers

import (
	"fmt"
	"strings"

	"github.com/joshuarubin/gil"
)

// LinkedListNode is a node in a linked list
type LinkedListNode struct {
	Value gil.Interface
	Next  *LinkedListNode
}

// LinkedList is an object that contains the Head pointer of a linked list
// as well as methods that apply to the whole list
type LinkedList struct {
	Head *LinkedListNode
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
		if node.Value != nil {
			ret = fmt.Sprintf("%s %s", ret, node.Value.String())
		} else if node.Next != nil {
			ret = fmt.Sprintf("%s nil", ret)
		}

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
