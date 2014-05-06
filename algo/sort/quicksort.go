package sort

import "github.com/joshuarubin/gil"

func qsSetPivot(work *qsWork) error {
	l := len(work.slice)

	// if the slice has only 1 or 2 elements, it doesn't matter which index is picked
	if l < 3 {
		work.pivot = l / 2
		return nil
	}

	// choose index of element with the median value of first, middle and last elements
	mid, last := l/2, l-1
	shortSlice := gil.Slice{work.slice[0], work.slice[mid], work.slice[last]}
	if err := NSmallest(shortSlice, 2); err != nil {
		return err
	}

	// median is the larger of the 2 smallest, so the 2nd value (index 1)
	switch shortSlice[1] {
	case work.slice[mid]:
		work.pivot = mid
	case work.slice[last]:
		work.pivot = last
	}

	return nil
}

type qsWork struct {
	slice gil.Slice
	pivot int
}

func (w *qsWork) Slice() gil.Slice {
	return w.slice
}

func (w *qsWork) SetSlice(slice gil.Slice) {
	w.slice = slice
}

func qsPartition(work *qsWork) (int, error) {
	valPivot := work.slice[work.pivot]

	last := len(work.slice) - 1

	// swap pivot and last values
	work.slice[work.pivot], work.slice[last] = work.slice[last], work.slice[work.pivot]

	store := 0
	for i, val := range work.slice {
		if val.Less(valPivot) {
			// swap slice[i] and slice[store]
			work.slice[i], work.slice[store] = work.slice[store], work.slice[i]
			store++
		}
	}

	// swap slice[store] and slice[right]
	// move pivot into its final place
	work.slice[store], work.slice[last] = work.slice[last], work.slice[store]

	return store, nil
}

func qs(workCh <-chan sortWork, resultCh chan<- error) {
	work, ok := (<-workCh).(*qsWork)
	if !ok {
		resultCh <- gil.TypeAssertionError{}
		return
	}

	if isShortSlice(work, false) {
		resultCh <- nil
		return
	}

	// choose pivot index
	if err := qsSetPivot(work); err != nil {
		resultCh <- err
		return
	}

	// do in-place less/greater partitioning
	pivotNew, err := qsPartition(work)
	if err != nil {
		resultCh <- err
		return
	}

	goWorkCh, goResultCh := make(chan sortWork), make(chan error)

	parts := []gil.Slice{
		work.slice[:pivotNew],
		work.slice[pivotNew+1:],
	}

	// sort each partition in its own goroutine
	for _, part := range parts {
		go qs(goWorkCh, goResultCh)
		goWorkCh <- &qsWork{slice: part}
	}

	// wait for results
	for _ = range parts {
		if err := <-goResultCh; err != nil {
			resultCh <- err
			return
		}
	}

	// and we're done
	resultCh <- nil
}

// Quick implements a generic, in-place, concurrent sort utilizing the
// quick sort algorithm. Pivot points are chosen as the index of the median
// value of the first, middle and last elements.
func Quick(slice gil.Slice) error {
	workCh, resultCh := make(chan sortWork), make(chan error)

	go qs(workCh, resultCh)
	workCh <- &qsWork{slice: slice}

	return <-resultCh
}
