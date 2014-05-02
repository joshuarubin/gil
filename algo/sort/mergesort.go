package sort

import "github.com/joshuarubin/gil"

func mergeSortMerge(left, right []gil.Interface) ([]gil.Interface, error) {
	llen, rlen := len(left), len(right)
	tlen := llen + rlen

	ret := make([]gil.Interface, tlen)

	i, j := 0, 0

	for k := 0; k < tlen; k++ {
		lrem, rrem := i < llen, j < rlen

		if lrem && rrem {
			// elements remain on both sides
			l, r := left[i], right[j]

			if less, err := l.Less(r); err != nil {
				return nil, err
			} else if less {
				ret[k] = l
				i++
			} else {
				ret[k] = r
				j++
			}
		} else if lrem {
			// elements only on left side
			ret[k] = left[i]
			i++
		} else {
			// elements only on right side
			ret[k] = right[j]
			j++
		}
	}

	return ret, nil
}

func MergeSort(list []gil.Interface) ([]gil.Interface, error) {
	l := len(list)
	if l < 2 {
		return list, nil
	}

	half := l / 2
	left, right := list[:half], list[half:]

	left, err := MergeSort(left)
	if err != nil {
		return nil, err
	}

	right, err = MergeSort(right)
	if err != nil {
		return nil, err
	}

	return mergeSortMerge(left, right)
}
