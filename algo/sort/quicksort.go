package sort

import "github.com/joshuarubin/gil"

func qsPivotIndex(list []gil.Interface) (int, error) {
	l := len(list)

	if l == 0 {
		return 0, gil.EmptyError("could not identify pivot of empty list")
	}

	// if the list has only 1 or 2 elements, it doesn't matter which index is picked
	if l < 3 {
		return l / 2, nil
	}

	// choose index of element with the median value of first, middle and last elements
	mid, last := l/2, l-1
	shortList := []gil.Interface{list[0], list[mid], list[last]}
	if err := NSmallest(shortList, 2); err != nil {
		return 0, err
	}

	// median is the larger of the 2 smallest, so the 2nd value (index 1)
	switch shortList[1] {
	case list[mid]:
		return mid, nil
	case list[last]:
		return last, nil
	}

	return 0, nil
}

func QuickSort(list []gil.Interface) ([]gil.Interface, error) {
	l := len(list)
	if l < 2 {
		ret := make([]gil.Interface, l)
		copy(ret, list)
		return ret, nil
	}

	iPivot, err := qsPivotIndex(list)
	valPivot := list[iPivot]

	if err != nil {
		return nil, err
	}

	low, high := make([]gil.Interface, l), make([]gil.Interface, l)
	j, k := 0, 0

	for i, val := range list {
		if i == iPivot {
			continue
		}

		if less, err := val.Less(valPivot); err != nil {
			return nil, err
		} else if less {
			low[j] = val
			j++
		} else {
			high[k] = val
			k++
		}
	}

	lowSorted, err := QuickSort(low[:j])
	if err != nil {
		return nil, err
	}

	highSorted, err := QuickSort(high[:k])
	if err != nil {
		return nil, err
	}

	copy(low, lowSorted)
	low[j] = valPivot
	copy(low[j+1:], highSorted)

	return low, nil
}
