package search

import "github.com/joshuarubin/gil"

func Binary(list []gil.Interface, val gil.Interface) (int, error) {
	l := len(list)

	valLess, err := val.Less(list[0])

	if err != nil {
		return 0, nil
	}

	testLess, err := list[l-1].Less(val)

	if err != nil {
		return 0, nil
	}

	if valLess || testLess {
		return 0, gil.RangeError{val}
	}

	var start, stop int = 0, l - 1

	for start <= stop {
		half := (stop + start) / 2
		test := list[half]

		if val == test {
			return half, nil
		}

		less, err := val.Less(test)

		if err != nil {
			return 0, err
		}

		if less {
			stop = half - 1
		} else {
			start = half + 1
		}
	}

	return 0, gil.NotFound{val}
}
