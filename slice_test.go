package gil

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	. "github.com/smartystreets/goconvey/convey"
)

type Result struct {
	Pos   int
	Found bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSliceFind(t *testing.T) {
	slice := CopyToIntSlice([]int{1, 3, 5, 6, 7, 9})
	expect := map[Int]Result{
		0:  {Pos: 0, Found: false},
		1:  {Pos: 0, Found: true},
		2:  {Pos: 1, Found: false},
		3:  {Pos: 1, Found: true},
		5:  {Pos: 2, Found: true},
		8:  {Pos: 5, Found: false},
		9:  {Pos: 5, Found: true},
		10: {Pos: 6, Found: false},
	}

	Convey("Given this slice", t, func() {
		for v, r := range expect {
			val, res := v, r // make variables local to this loop only
			pos := slice.Find(val)

			if res.Found == true {
				Convey(fmt.Sprintf("The position of %d should be %d", val, res.Pos), func() {
					So(pos, ShouldEqual, res.Pos)
				})

				Convey(fmt.Sprintf("No error should be returned when searching for %d", val), func() {
					So(pos, ShouldBeLessThan, len(slice))
					So(slice[pos], ShouldEqual, val)
				})

				continue
			}

			Convey(fmt.Sprintf("%d should not be found", val), func() {
				So(pos == len(slice) || slice[pos] != val, ShouldBeTrue)
			})
		}
	})

	Convey("Given a non-homogenous slice", t, func() {
		Convey("Range checks should fail", func() {
			slice := Slice{
				Int(0),
				String("a"),
			}

			pos := slice.Find(String("z"))
			So(pos, ShouldEqual, 2)

			pos = slice.Find(Int(4))
			So(pos, ShouldEqual, 2)
		})

		Convey("Comparison checks should fail", func() {
			slice := Slice{
				Int(0),
				String("a"),
				Int(9),
			}

			pos := slice.Find(String("z"))
			So(pos, ShouldEqual, 3)

			pos = slice.Find(Int(4))
			So(pos, ShouldEqual, 2)
			So(slice[pos], ShouldNotEqual, 4)
		})
	})
}

func testSorted(t *testing.T, slice, sorted Slice) {
	Convey(fmt.Sprintf("For sorting a random int slice (size %d)", len(slice)), t, func() {
		Convey("The length should not change", func() {
			So(len(sorted), ShouldEqual, len(slice))
		})

		if len(slice) > 0 {
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

func genSlice(size int) Slice {
	slice := make(Slice, size)

	for i := 0; i < size; i++ {
		slice[i] = Int(rand.Int())
	}

	return slice
}

func benchmarkSlice(size int) Slice {
	slice := make(Slice, size)
	for i := 0; i < size; i++ {
		slice[i] = Int(rand.Int())
	}
	return slice
}

func BenchmarkSort(b *testing.B) {
	slice := benchmarkSlice(2 ^ 14)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copyOfSlice := make(Slice, len(slice))
		copy(copyOfSlice, slice)
		b.StartTimer()

		copyOfSlice.Sort()
	}
}

func TestSliceSort(t *testing.T) {
	for _, size := range []int{0, 1, 2, 3, 10, 100} {
		slice := genSlice(size)
		sorted := make(Slice, size)
		copy(sorted, slice)
		sorted.Sort()
		testSorted(t, slice, sorted)
	}
}
