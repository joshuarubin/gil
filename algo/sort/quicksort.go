package sort

import "github.com/joshuarubin/gil"

func qsSetPivot(work *qsWork) error {
	l := len(work.list)

	// if the list has only 1 or 2 elements, it doesn't matter which index is picked
	if l < 3 {
		work.pivot = l / 2
		return nil
	}

	// choose index of element with the median value of first, middle and last elements
	mid, last := l/2, l-1
	shortList := []gil.Interface{work.list[0], work.list[mid], work.list[last]}
	if err := NSmallest(shortList, 2); err != nil {
		return err
	}

	// median is the larger of the 2 smallest, so the 2nd value (index 1)
	switch shortList[1] {
	case work.list[mid]:
		work.pivot = mid
	case work.list[last]:
		work.pivot = last
	}

	return nil
}

type qsWork struct {
	list  []gil.Interface
	pivot int
}

func (w *qsWork) List() []gil.Interface {
	return w.list
}

func (w *qsWork) SetList(list []gil.Interface) {
	w.list = list
}

func qsPartition(work *qsWork) (int, error) {
	valPivot := work.list[work.pivot]

	last := len(work.list) - 1

	// swap pivot and last values
	work.list[work.pivot], work.list[last] = work.list[last], work.list[work.pivot]

	store := 0
	for i, val := range work.list {
		if less, err := val.Less(valPivot); err != nil {
			return 0, err
		} else if less {
			// swap list[i] and list[store]
			work.list[i], work.list[store] = work.list[store], work.list[i]
			store++
		}
	}

	// swap list[store] and list[right]
	// move pivot into its final place
	work.list[store], work.list[last] = work.list[last], work.list[store]

	return store, nil
}

func qs(workCh <-chan sortWork, resultCh chan<- error) {
	work, ok := (<-workCh).(*qsWork)
	if !ok {
		resultCh <- gil.TypeAssertionError{}
		return
	}

	if isShortList(work, false) {
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

	parts := [][]gil.Interface{
		work.list[:pivotNew],
		work.list[pivotNew+1:],
	}

	// sort each partition in its own goroutine
	for _, part := range parts {
		go qs(goWorkCh, goResultCh)
		goWorkCh <- &qsWork{list: part}
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
func Quick(list []gil.Interface) error {
	workCh, resultCh := make(chan sortWork), make(chan error)

	go qs(workCh, resultCh)
	workCh <- &qsWork{list: list}

	return <-resultCh
}
