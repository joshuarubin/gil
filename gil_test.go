package gil

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {
	Convey("Errors should work properly", t, func() {
		So(NotFoundError{Int(0)}.Error(), ShouldEqual, "not found")
		So(TypeAssertionError{}.Error(), ShouldEqual, "type assertion error")
		So(ArgumentError("some error text").Error(), ShouldEqual, "some error text")
	})
}

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

type Result struct {
	Pos int
	Err error
}

func TestFind(t *testing.T) {
	slice := CopyToIntSlice([]int{1, 3, 5, 6, 7, 9})
	expect := map[Int]Result{
		0:  {Pos: 0, Err: NotFoundError{Interface: Int(0)}},
		1:  {Pos: 0, Err: nil},
		2:  {Pos: 1, Err: NotFoundError{Interface: Int(2)}},
		3:  {Pos: 1, Err: nil},
		5:  {Pos: 2, Err: nil},
		8:  {Pos: 5, Err: NotFoundError{Interface: Int(8)}},
		9:  {Pos: 5, Err: nil},
		10: {Pos: 6, Err: NotFoundError{Interface: Int(10)}},
	}

	Convey("Given this slice", t, func() {
		for v, r := range expect {
			val, res := v, r // make variables local to this loop only
			pos, err := slice.Find(val)

			if res.Err == nil {
				Convey(fmt.Sprintf("The position of %d should be %d", val, res.Pos), func() {
					So(pos, ShouldEqual, res.Pos)
				})

				Convey(fmt.Sprintf("No error should be returned when searching for %d", val), func() {
					So(err, ShouldBeNil)
				})

				continue
			}

			Convey(fmt.Sprintf("%d should not be found", val), func() {
				So(err, ShouldNotBeNil)
			})

			Convey(fmt.Sprintf("Error Type %T should be returned when searching for %d", res.Err, val), func() {
				So(err, ShouldHaveSameTypeAs, res.Err)
			})
		}
	})

	Convey("Given a non-homogenous slice", t, func() {
		Convey("Range checks should fail", func() {
			slice := Slice{
				Int(0),
				String("a"),
			}

			pos, err := slice.Find(String("z"))

			So(pos, ShouldEqual, 2)
			So(err, ShouldHaveSameTypeAs, NotFoundError{})

			pos, err = slice.Find(Int(4))

			So(pos, ShouldEqual, 2)
			So(err, ShouldHaveSameTypeAs, NotFoundError{})
		})

		Convey("Comparison checks should fail", func() {
			slice := Slice{
				Int(0),
				String("a"),
				Int(9),
			}

			pos, err := slice.Find(String("z"))

			So(pos, ShouldEqual, 3)
			So(err, ShouldHaveSameTypeAs, NotFoundError{})

			pos, err = slice.Find(Int(4))

			So(pos, ShouldEqual, 2)
			So(err, ShouldHaveSameTypeAs, NotFoundError{})
		})
	})

}
