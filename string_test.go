package gil

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	Convey("String should work", t, func() {
		slice := CopyToStringSlice([]string{"a", "b", "c", "d", "e"})

		prev := slice[0]
		for i := 1; i < len(slice); i++ {
			cur := slice[i]

			str, ok := cur.(String)
			So(ok, ShouldBeTrue)
			So(len(str), ShouldEqual, 1)
			So(str[0], ShouldEqual, 'a'+i)

			So(prev.Less(cur), ShouldBeTrue)
			So(cur.Less(prev), ShouldBeFalse)
			So(prev == cur, ShouldBeFalse)

			So(prev.Less(Int(1)), ShouldBeFalse)
			So(cur.Less(Int(1)), ShouldBeFalse)

			prev = cur
		}
	})
}
