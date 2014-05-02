package sort

import "github.com/joshuarubin/gil"

// puts the n smallest values at the beginning of list, in order
func NSmallest(list []gil.Interface, n int) error {
	// modifies list
	// uses a partial selection sort
	for i, val := range list[:n-1] {
		iMin := i
		valMin := val

		for j, next := range list[i:] {
			if less, err := next.Less(valMin); err != nil {
				return err
			} else if less {
				iMin = j
				valMin = next
			}
		}

		list[i], list[iMin] = list[iMin], list[i] // swap
	}

	return nil
}
