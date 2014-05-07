package gil

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	Convey("Int should work", t, func() {
		slice := CopyToIntSlice([]int{0, 1, 2, 3, 4, 5})

		prev := slice[0]
		for i := 1; i < len(slice); i++ {
			cur := slice[i]

			So(cur, ShouldEqual, i)
			So(prev.Less(cur), ShouldBeTrue)
			So(cur.Less(prev), ShouldBeFalse)
			So(prev == cur, ShouldBeFalse)

			So(prev.Less(String("a")), ShouldBeFalse)
			So(cur.Less(String("a")), ShouldBeFalse)

			prev = cur
		}
	})
}
