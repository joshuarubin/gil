package containers

import (
	"fmt"
	"testing"

	"github.com/joshuarubin/gil"
	. "github.com/smartystreets/goconvey/convey"
)

const NUM = 10

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
