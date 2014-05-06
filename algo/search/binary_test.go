package search

import (
	"fmt"
	"testing"

	"github.com/joshuarubin/gil"
	. "github.com/smartystreets/goconvey/convey"
)

type Result struct {
	Pos int
	Err error
}

func TestBinarySearch(t *testing.T) {
	slice := gil.CopyToIntSlice([]int{1, 3, 5, 6, 7, 9})
	expect := map[gil.Int]Result{
		0:  {Pos: 0, Err: gil.NotFoundError{Interface: gil.Int(0)}},
		1:  {Pos: 0, Err: nil},
		2:  {Pos: 1, Err: gil.NotFoundError{Interface: gil.Int(2)}},
		3:  {Pos: 1, Err: nil},
		5:  {Pos: 2, Err: nil},
		8:  {Pos: 5, Err: gil.NotFoundError{Interface: gil.Int(8)}},
		9:  {Pos: 5, Err: nil},
		10: {Pos: 6, Err: gil.NotFoundError{Interface: gil.Int(10)}},
	}

	Convey("Given this slice", t, func() {
		for v, r := range expect {
			val, res := v, r // make variables local to this loop only
			pos, err := BinarySearchSlice(slice, val)

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
			slice := []gil.Interface{
				gil.Int(0),
				gil.String("a"),
			}

			pos, err := BinarySearchSlice(slice, gil.String("z"))

			So(pos, ShouldEqual, 2)
			So(err, ShouldHaveSameTypeAs, gil.NotFoundError{})

			pos, err = BinarySearchSlice(slice, gil.Int(4))

			So(pos, ShouldEqual, 2)
			So(err, ShouldHaveSameTypeAs, gil.NotFoundError{})
		})

		Convey("Comparison checks should fail", func() {
			slice := []gil.Interface{
				gil.Int(0),
				gil.String("a"),
				gil.Int(9),
			}

			pos, err := BinarySearchSlice(slice, gil.String("z"))

			So(pos, ShouldEqual, 3)
			So(err, ShouldHaveSameTypeAs, gil.NotFoundError{})

			pos, err = BinarySearchSlice(slice, gil.Int(4))

			So(pos, ShouldEqual, 2)
			So(err, ShouldHaveSameTypeAs, gil.NotFoundError{})
		})
	})
}
