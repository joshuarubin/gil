package search

import "github.com/joshuarubin/gil"

// Binary implements a generic binary search. The index of the matching value
// is returned. If no match is found, NotFoundError is returned.
func Binary(list []gil.Interface, val gil.Interface) (int, error) {
	l := len(list)

	if val.Less(list[0]) || list[l-1].Less(val) {
		return 0, gil.RangeError{Interface: val}
	}

	var start, stop int = 0, l - 1

	for start <= stop {
		half := (stop + start) / 2
		test := list[half]

		if val == test {
			return half, nil
		}

		if val.Less(test) {
			stop = half - 1
		} else {
			start = half + 1
		}
	}

	return 0, gil.NotFoundError{Interface: val}
}
