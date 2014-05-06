package sort

import "github.com/joshuarubin/gil"

// TODO(jrubin) test this

// NSmallest implements a partial selection sort. It modifies, in-place,
// the passed in slice such that the n smallest values are placed at the
// beginning of the slice, in order.
func NSmallest(slice gil.Slice, n int) error {
	// modifies slice
	// uses a partial selection sort
	for i, val := range slice[:n-1] {
		iMin := i
		valMin := val

		for j, next := range slice[i:] {
			if next.Less(valMin) {
				iMin = j
				valMin = next
			}
		}

		slice[i], slice[iMin] = slice[iMin], slice[i] // swap
	}

	return nil
}
