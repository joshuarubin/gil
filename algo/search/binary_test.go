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
	list := gil.CopyToIntSlice([]int{1, 3, 5, 6, 7, 9})
	expect := map[gil.Int]Result{
		0:  {0, gil.RangeError{Interface: gil.Int(0)}},
		1:  {0, nil},
		2:  {0, gil.NotFoundError{Interface: gil.Int(2)}},
		3:  {1, nil},
		5:  {2, nil},
		8:  {0, gil.NotFoundError{Interface: gil.Int(2)}},
		9:  {5, nil},
		10: {0, gil.RangeError{Interface: gil.Int(0)}},
	}

	Convey("Given this list", t, func() {
		for v, r := range expect {
			val, res := v, r // make variables local to this loop only
			pos, err := Binary(list, val)

			if res.Err == nil {
				Convey(fmt.Sprintf("The position of %d should be %d", val, res.Pos), func() {
					So(res.Pos, ShouldEqual, pos)
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

	Convey("Given a non-homogenous list", t, func() {
		Convey("Range checks should fail", func() {
			list := []gil.Interface{
				gil.Int(0),
				gil.String("a"),
			}

			pos, err := Binary(list, gil.String("z"))

			So(pos, ShouldEqual, 0)
			So(err, ShouldHaveSameTypeAs, gil.TypeMismatchError{})

			pos, err = Binary(list, gil.Int(4))

			So(pos, ShouldEqual, 0)
			So(err, ShouldHaveSameTypeAs, gil.TypeMismatchError{})
		})

		Convey("Comparison checks should fail", func() {
			list := []gil.Interface{
				gil.Int(0),
				gil.String("a"),
				gil.Int(9),
			}

			pos, err := Binary(list, gil.String("z"))

			So(pos, ShouldEqual, 0)
			So(err, ShouldHaveSameTypeAs, gil.TypeMismatchError{})

			pos, err = Binary(list, gil.Int(4))

			So(pos, ShouldEqual, 0)
			So(err, ShouldHaveSameTypeAs, gil.TypeMismatchError{})
		})
	})
}
