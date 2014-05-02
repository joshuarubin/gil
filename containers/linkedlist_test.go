package containers

import (
	"testing"

	"github.com/joshuarubin/gil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkedListReverse(t *testing.T) {
	Convey("Linked list should be reversed", t, func() {
		list := &LinkedList{&LinkedListNode{}}

		const NUM = 10

		list.Walk(func(node *LinkedListNode, i int) bool {
			node.Value = gil.Int(i)

			if i < NUM {
				node.Next = &LinkedListNode{}
				return true
			}

			return false
		})

		list.Reverse()

		list.Walk(func(node *LinkedListNode, i int) bool {
			So(node.Value, ShouldEqual, NUM-i)
			return true
		})
	})
}
