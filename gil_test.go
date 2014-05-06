package gil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {
	Convey("Errors should work properly", t, func() {
		So(RangeError{Int(0)}.Error(), ShouldEqual, "not in range")
		So(NotFoundError{Int(0)}.Error(), ShouldEqual, "not found")
		So(TypeAssertionError{}.Error(), ShouldEqual, "type assertion error")
		So(ArgumentError("some error text").Error(), ShouldEqual, "some error text")
	})
}

func TestInt(t *testing.T) {
	Convey("Int should work", t, func() {
		list := CopyToIntSlice([]int{0, 1, 2, 3, 4, 5})

		prev := list[0]
		for i := 1; i < len(list); i++ {
			cur := list[i]

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

func TestString(t *testing.T) {
	Convey("String should work", t, func() {
		list := CopyToStringSlice([]string{"a", "b", "c", "d", "e"})

		prev := list[0]
		for i := 1; i < len(list); i++ {
			cur := list[i]

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
