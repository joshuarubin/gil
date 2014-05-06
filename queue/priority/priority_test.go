package priority

import (
	"math/rand"
	"testing"
	"time"

	"github.com/joshuarubin/gil"
	. "github.com/smartystreets/goconvey/convey"
)

const NUM = 10

func TestPriorityQueue(t *testing.T) {
	q := New()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < NUM; i++ {
		q.Push(gil.Int(rand.Int()))
	}

	prev := q.Peek()

	Convey("PriorityQueue should implement a priority queue", t, func() {
		for i := 0; i < NUM; i++ {
			So(q.Len(), ShouldEqual, NUM-i)

			cur := q.Pop()
			So(q.Len(), ShouldEqual, NUM-i-1)
			So(prev.Less(cur) || cur == prev, ShouldBeTrue)

			prev = cur
		}
	})
}
