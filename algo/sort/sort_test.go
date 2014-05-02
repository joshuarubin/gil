package sort

import (
	"testing"

	"github.com/joshuarubin/gil"
)

var list = gil.CopyToIntSlice([]int{
	62, 34, 10, 27, 62,
	24, 11, 99, 71, 71,
	45, 83, 71, 18, 29,
	62, 8, 54, 3, 41,
	91, 42, 1, 74, 7,
	81, 14, 73, 56, 47,
	19, 78, 65, 10, 35,
})

func TestMergeSort(t *testing.T) {
	sorted, err := MergeSort(list)
	testSorted(t, sorted, err)
}

func TestQuickSort(t *testing.T) {
	sorted, err := QuickSort(list)
	testSorted(t, sorted, err)
}

func testSorted(t *testing.T, sorted []gil.Interface, err error) {
	if err != nil {
		t.Fatal("unexpected sort error")
	}

	if len(sorted) != len(list) {
		t.Fatal("list length changed")
	}

	prev := sorted[0]
	for i, val := range sorted[1:] {
		less, err := val.Less(prev)

		if err != nil {
			t.Fatal("unexpected comparison error")
		} else if less {
			t.Errorf("%d < %d at index %d", val, prev, i)
		}

		prev = val
	}

	if !containsAllOf(list, sorted) {
		t.Error("sorted slice is missing values")
	}

	if !containsAllOf(sorted, list) {
		t.Error("sorted slice has extra values")
	}
}

func contains(test []gil.Interface, val gil.Interface) bool {
	l := len(test)
	for i := 0; i < l; i++ {
		if test[i] == val {
			return true
		}
	}
	return false
}

func containsAllOf(match, test []gil.Interface) bool {
	for _, val := range match {
		if !contains(test, val) {
			return false
		}
	}

	return true
}
