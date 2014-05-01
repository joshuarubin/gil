package sort

import "github.com/joshuarubin/gil"

func merge(left, right []gil.Interface) ([]gil.Interface, error) {
	llen, rlen := len(left), len(right)
	tlen := llen + rlen

	ret := make([]gil.Interface, tlen)

	i, j := 0, 0

	for k := 0; k < tlen; k++ {
		lrem, rrem := i < llen, j < rlen

		if lrem && rrem {
			// elements remain on both sides
			l, r := left[i], right[j]

			less, err := l.Less(r)

			if err != nil {
				return ret, err
			}

			if less {
				//fmt.Println(k, l, "<")
				ret[k] = l
				i++
			} else {
				//fmt.Println(k, r, ">")
				ret[k] = r
				j++
			}
		} else if lrem {
			//fmt.Println(k, left[i], "lrem")
			// elements only on left side
			ret[k] = left[i]
			i++
		} else {
			//fmt.Println(k, right[j], "rrem")
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

	var err error
	left, err = MergeSort(left)
	if err != nil {
		return list, err
	}

	right, err = MergeSort(right)
	if err != nil {
		return list, err
	}

	return merge(left, right)
}
