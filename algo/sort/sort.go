package msort

func merge(left, right []int) []int {
	llen := len(left)
	rlen := len(right)
	tlen := llen + rlen

	ret := make([]int, tlen)

	i, j := 0, 0

	for k := 0; k < tlen; k++ {
		lrem := i < llen
		rrem := j < rlen

		if lrem && rrem {
			// elements remain on both sides
			l := left[i]
			r := right[j]

			if l < r {
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

	return ret
}

func MSort(list []int) {
	l := len(list)
	if l < 2 {
		return
	}

	half := l / 2
	left := list[:half]
	right := list[half:]

	MSort(left)
	MSort(right)

	copy(list, merge(left, right))
}
