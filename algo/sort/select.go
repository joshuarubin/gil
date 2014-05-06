package sort

import "github.com/joshuarubin/gil"

// TODO(jrubin) test this

// NSmallest implements a partial selection sort. It modifies, in-place,
// the passed in list such that the n smallest values are placed at the
// beginning of the list, in order.
func NSmallest(list []gil.Interface, n int) error {
	// modifies list
	// uses a partial selection sort
	for i, val := range list[:n-1] {
		iMin := i
		valMin := val

		for j, next := range list[i:] {
			if next.Less(valMin) {
				iMin = j
				valMin = next
			}
		}

		list[i], list[iMin] = list[iMin], list[i] // swap
	}

	return nil
}
