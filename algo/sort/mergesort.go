package sort

import "github.com/joshuarubin/gil"

type mergeSortData struct {
	pos  int
	list []gil.Interface
}

func (m *mergeSortData) Peek() gil.Interface {
	return m.list[m.pos]
}

func (m *mergeSortData) Pop() gil.Interface {
	ret := m.Peek()
	m.pos++
	return ret
}

func mergeSortInterleave(parts ...[]gil.Interface) ([]gil.Interface, error) {
	const NUM = 2

	if len(parts) != NUM {
		return nil, gil.ArgumentError("wrong number of parts")
	}

	data := [NUM]*mergeSortData{}
	tlen := 0
	for side, part := range parts {
		data[side] = &mergeSortData{list: part}
		tlen += len(part)
	}

	ret := make([]gil.Interface, tlen)

	for k := 0; k < tlen; k++ {
		rem := 0
		for side, d := range data {
			if d.pos < len(d.list) {
				rem |= 1 << uint(side)
			}
		}

		if rem == 3 { // 0b11
			// elements remain on both sides
			var min *mergeSortData
			for _, d := range data {
				if min == nil {
					min = d
					continue
				}

				if less, err := d.Peek().Less(min.Peek()); err != nil {
					return nil, err
				} else if less {
					min = d
				}
			}

			ret[k] = min.Pop()
			continue
		}

		// data remains on only one side
		index := -1
		for rem > 0 {
			rem >>= 1
			index++
		}
		ret[k] = data[index].Pop()
	}

	return ret, nil
}

type mergeSortWork struct {
	list []gil.Interface
	side int
}

type mergeSortResult struct {
	mergeSortWork
	err error
}

func mergeSortConcurrent(workCh <-chan mergeSortWork, resultCh chan<- mergeSortResult) {
	work := <-workCh

	if len(work.list) < 2 {
		resultCh <- mergeSortResult{work, nil}
		return
	}

	half := len(work.list) / 2
	parts := [][]gil.Interface{
		work.list[:half],
		work.list[half:],
	}

	goWorkCh, goResultCh := make(chan mergeSortWork), make(chan mergeSortResult)

	for side, part := range parts {
		go mergeSortConcurrent(goWorkCh, goResultCh)
		goWorkCh <- mergeSortWork{part, side}
	}

	for _ = range parts {
		result := <-goResultCh

		if result.err != nil {
			resultCh <- mergeSortResult{err: result.err}
			return
		}

		parts[result.side] = result.list
	}

	if result, err := mergeSortInterleave(parts[0], parts[1]); err != nil {
		resultCh <- mergeSortResult{err: err}
	} else {
		resultCh <- mergeSortResult{mergeSortWork{result, work.side}, nil}
	}
}

func MergeSort(list []gil.Interface) ([]gil.Interface, error) {
	workCh, resultCh := make(chan mergeSortWork), make(chan mergeSortResult)
	go mergeSortConcurrent(workCh, resultCh)
	workCh <- mergeSortWork{list: list}
	result := <-resultCh
	return result.list, result.err
}
