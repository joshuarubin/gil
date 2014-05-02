package sort

import (
	"fmt"
	"testing"

	"github.com/joshuarubin/gil"
	. "github.com/smartystreets/goconvey/convey"
)

func genList(size int) []gil.Interface {
	if size < 0 {
		return nil
	}

	list := gil.CopyToIntSlice([]int{
		62, 34, 10, 27, 62,
		24, 11, 99, 71, 71,
		45, 83, 71, 18, 29,
		62, 8, 54, 3, 41,
		91, 42, 1, 74, 7,
		81, 14, 73, 56, 47,
		19, 78, 65, 10, 35,
	})

	if size <= len(list) {
		return list[:size]
	}

	return list
}

func TestMergeSort(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 2, 3, 10, 100} {
		list := genList(size)
		sorted, err := Merge(list)
		testSorted(t, "Merge", size, list, sorted, err)
	}
}

func TestQuickSort(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 2, 3, 10, 100} {
		list := genList(size)
		sorted := make([]gil.Interface, len(list))
		copy(sorted, list)
		err := Quick(sorted)
		testSorted(t, "Quick", size, list, sorted, err)
	}
}

func testSorted(t *testing.T, algo string, size int, list, sorted []gil.Interface, err error) {
	Convey(fmt.Sprintf("For the %sSort algorithm (size %d)", algo, size), t, func() {
		Convey("There should be no error", func() {
			So(err, ShouldBeNil)
		})

		Convey("Original list should not be modified", func() {
			So(list, ShouldResemble, genList(size))
		})

		Convey("The length should not change", func() {
			So(len(sorted), ShouldEqual, len(list))
		})

		if size > 0 {
			Convey("The values should be in order", func() {
				prev := sorted[0]
				for _, val := range sorted[1:] {
					less, err := val.Less(prev)

					So(err, ShouldBeNil)
					So(less, ShouldBeFalse)

					prev = val
				}
			})
		}

		Convey("No values should be missing", func() {
			for _, val := range list {
				So(sorted, ShouldContain, val)
			}
		})

		Convey("No values should be added", func() {
			for _, val := range sorted {
				So(list, ShouldContain, val)
			}
		})
	})
}
