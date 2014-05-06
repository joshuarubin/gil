package sort

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joshuarubin/gil"
	. "github.com/smartystreets/goconvey/convey"
)

func genSlice(size int) gil.Slice {
	if size < 0 {
		return nil
	}

	slice := gil.CopyToIntSlice([]int{
		62, 34, 10, 27, 62,
		24, 11, 99, 71, 71,
		45, 83, 71, 18, 29,
		62, 8, 54, 3, 41,
		91, 42, 1, 74, 7,
		81, 14, 73, 56, 47,
		19, 78, 65, 10, 35,
	})

	if size <= len(slice) {
		return slice[:size]
	}

	return slice
}

func TestMergeSort(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 2, 3, 10, 100} {
		slice := genSlice(size)
		sorted, err := Merge(slice)
		testSorted(t, "Merge", size, slice, sorted, err)
	}
}

func TestQuickSort(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 2, 3, 10, 100} {
		slice := genSlice(size)
		sorted := make(gil.Slice, len(slice))
		copy(sorted, slice)
		err := Quick(sorted)
		testSorted(t, "Quick", size, slice, sorted, err)
	}
}

func testSorted(t *testing.T, algo string, size int, slice, sorted gil.Slice, err error) {
	Convey(fmt.Sprintf("For the %sSort algorithm (size %d)", algo, size), t, func() {
		Convey("There should be no error", func() {
			So(err, ShouldBeNil)
		})

		Convey("Original slice should not be modified", func() {
			So(slice, ShouldResemble, genSlice(size))
		})

		Convey("The length should not change", func() {
			So(len(sorted), ShouldEqual, len(slice))
		})

		if size > 0 {
			Convey("The values should be in order", func() {
				prev := sorted[0]
				for _, val := range sorted[1:] {
					So(val.Less(prev), ShouldBeFalse)
					prev = val
				}
			})
		}

		Convey("No values should be missing", func() {
			for _, val := range slice {
				So(sorted, ShouldContain, val)
			}
		})

		Convey("No values should be added", func() {
			for _, val := range sorted {
				So(slice, ShouldContain, val)
			}
		})
	})
}

func benchmarkSlice(size int) gil.Slice {
	slice := make(gil.Slice, size)
	for i := 0; i < size; i++ {
		slice[i] = gil.Int(rand.Int())
	}
	return slice
}

func BenchmarkMergeSort(b *testing.B) {
	slice := benchmarkSlice(2 ^ 14)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copyOfSlice := make(gil.Slice, len(slice))
		copy(copyOfSlice, slice)
		b.StartTimer()

		Merge(copyOfSlice)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	slice := benchmarkSlice(2 ^ 14)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copyOfSlice := make(gil.Slice, len(slice))
		copy(copyOfSlice, slice)
		b.StartTimer()

		Quick(copyOfSlice)
	}
}
