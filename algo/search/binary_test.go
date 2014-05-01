package search

import (
	"testing"

	"github.com/joshuarubin/gil"
)

type Result struct {
	Pos int
	Err error
}

func TestBinarySearch(t *testing.T) {
	list := gil.CopyToIntSlice([]int{1, 3, 5, 6, 7, 9})
	expect := map[gil.Int]Result{
		0:  {0, gil.RangeError{gil.Int(0)}},
		1:  {0, nil},
		2:  {0, gil.NotFound{gil.Int(2)}},
		3:  {1, nil},
		5:  {2, nil},
		8:  {0, gil.NotFound{gil.Int(2)}},
		9:  {5, nil},
		10: {0, gil.RangeError{gil.Int(0)}},
	}

	for val, res := range expect {
		pos, err := Binary(list, val)
		if res.Pos != pos {
			t.Errorf("%d != %d", res.Pos, pos)
		}

		if res.Err == nil {
			if err != nil {
				t.Errorf("Unexpected error, not nil (%T), for value %d (expected %d)", err, val, res.Pos)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error, but didn't find one.")
			} else if _, ok := res.Err.(gil.RangeError); !ok {
				if _, ok := err.(gil.RangeError); ok {
					t.Errorf("Unexpected error, not RangeError")
				}
			} else if _, ok := res.Err.(gil.NotFound); !ok {
				if _, ok := err.(gil.NotFound); ok {
					t.Errorf("Unexpected error, not NotFound")
				}
			}
		}
	}
}
