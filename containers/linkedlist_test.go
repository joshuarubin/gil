package containers

import (
	"testing"

	"github.com/joshuarubin/gil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkedList(t *testing.T) {
	const NUM = 10

	genList := func() *LinkedList {
		list := &LinkedList{&LinkedListNode{}}

		list.Walk(func(node *LinkedListNode, i int) bool {
			node.Value = gil.Int(i)

			if i < NUM {
				node.Next = &LinkedListNode{}
				return true
			}

			return false
		})

		return list
	}

	Convey("Linked list should print the correct string", t, func() {
		list := genList()
		So(list.String(), ShouldEqual, "0 1 2 3 4 5 6 7 8 9 10")
	})

	Convey("Linked list should be reversed", t, func() {
		list := genList()
		list.Reverse()
		list.Walk(func(node *LinkedListNode, i int) bool {
			So(node.Value, ShouldEqual, NUM-i)
			return true
		})
	})
}
