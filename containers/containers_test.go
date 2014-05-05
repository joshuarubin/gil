package containers

import (
	"fmt"
	"testing"

	"github.com/joshuarubin/gil"
	. "github.com/smartystreets/goconvey/convey"
)

const NUM = 10

func TestLinkedList(t *testing.T) {
	genList := func() *LinkedList {
		list := &LinkedList{}
		for i := NUM - 1; i >= 0; i-- {
			list.Push(i)
		}
		return list
	}

	Convey("Linked list should print the correct string", t, func() {
		list := genList()
		So(list.String(), ShouldEqual, "0 1 2 3 4 5 6 7 8 9")
		list.Head.RemoveNext()
		So(list.String(), ShouldEqual, "0 2 3 4 5 6 7 8 9")

		for i := 0; i < NUM-1; i++ {
			list.Pop()
		}

		So(list.String(), ShouldEqual, "")
		So(list.Pop(), ShouldBeNil)

		list.Push(nil)
		So(list.String(), ShouldEqual, "<nil>")
	})

	Convey("Linked list should be reversed", t, func() {
		list := genList()
		list.Reverse()
		list.Walk(func(node *LinkedListNode, i int) bool {
			So(node.Value, ShouldEqual, NUM-1-i)
			return true
		})
	})

	Convey("List should be walked", t, func() {
		list := genList()

		i := 0
		list.Walk(func(node *LinkedListNode, num int) bool {
			if i == 5 {
				return false
			}

			So(i, ShouldEqual, num)
			So(i, ShouldBeLessThan, 5)

			i++
			return true
		})
	})
}

func testPopulateQueue(t *testing.T, q gil.Queue, qType string) {
	Convey(fmt.Sprintf("%s Queue should be populated", qType), t, func() {
		So(q.Len(), ShouldEqual, 0)
		So(q.Peek(), ShouldBeNil)
		So(q.Pop(), ShouldBeNil)

		for i := 0; i < NUM; i++ {
			q.Push(gil.Int(i))
			So(q.Len(), ShouldEqual, i+1)
		}

		So(q.Len(), ShouldEqual, NUM)
	})
}

func TestStack(t *testing.T) {
	stack := &Stack{}
	testPopulateQueue(t, stack, "LIFO")

	Convey("Stack should implement LIFO", t, func() {
		for i := 0; i < NUM; i++ {
			So(stack.Peek(), ShouldEqual, NUM-1-i)
			So(stack.Len(), ShouldEqual, NUM-i)
			So(stack.Pop(), ShouldEqual, NUM-1-i)
			So(stack.Len(), ShouldEqual, NUM-1-i)
		}
	})
}

func TestQueue(t *testing.T) {
	q := &Queue{}
	testPopulateQueue(t, q, "FIFO")

	Convey("Queue should implement FIFO", t, func() {
		for i := 0; i < NUM; i++ {
			So(q.Peek(), ShouldEqual, i)
			So(q.Len(), ShouldEqual, NUM-i)
			So(q.Pop(), ShouldEqual, i)
			So(q.Len(), ShouldEqual, NUM-1-i)
		}
	})
}
