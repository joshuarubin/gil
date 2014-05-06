package sort

import "github.com/joshuarubin/gil"

type msData struct {
	pos   int
	slice gil.Slice
}

func (m *msData) Peek() gil.Interface {
	return m.slice[m.pos]
}

func (m *msData) Pop() gil.Interface {
	ret := m.Peek()
	m.pos++
	return ret
}

func msInterleave(parts ...gil.Slice) (gil.Slice, error) {
	const NUM = 2

	if len(parts) != NUM {
		return nil, gil.ArgumentError("wrong number of parts")
	}

	data := [NUM]*msData{}
	tlen := 0
	for side, part := range parts {
		data[side] = &msData{slice: part}
		tlen += len(part)
	}

	ret := make(gil.Slice, tlen)

	for k := 0; k < tlen; k++ {
		rem := 0
		for side, d := range data {
			if d.pos < len(d.slice) {
				rem |= 1 << uint(side)
			}
		}

		if rem == 3 { // 0b11
			// elements remain on both sides
			var min *msData
			for _, d := range data {
				if min == nil {
					min = d
					continue
				}

				if d.Peek().Less(min.Peek()) {
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

type msWork struct {
	slice gil.Slice
	side  int
	err   error
}

func (w *msWork) Slice() gil.Slice {
	return w.slice
}

func (w *msWork) SetSlice(slice gil.Slice) {
	w.slice = slice
}

func isShortSlice(work sortWork, copySlice bool) bool {
	l := len(work.Slice())

	if l >= 2 {
		return false
	}

	if copySlice {
		ret := make(gil.Slice, l)
		copy(ret, work.Slice())
		work.SetSlice(ret)
	}

	return true
}

func ms(workCh <-chan sortWork, resultCh chan<- sortWork) {
	work, ok := (<-workCh).(*msWork)
	if !ok {
		resultCh <- &msWork{err: gil.TypeAssertionError{}}
		return
	}

	if isShortSlice(work, true) {
		resultCh <- work
		return
	}

	half := len(work.slice) / 2
	parts := []gil.Slice{
		work.slice[:half],
		work.slice[half:],
	}

	goWorkCh, goResultCh := make(chan sortWork), make(chan sortWork)

	for side, part := range parts {
		go ms(goWorkCh, goResultCh)
		goWorkCh <- &msWork{part, side, nil}
	}

	for _ = range parts {
		result, ok := (<-goResultCh).(*msWork)

		if !ok {
			resultCh <- &msWork{err: gil.TypeAssertionError{}}
			return
		}

		if result.err != nil {
			resultCh <- result
			return
		}

		parts[result.side] = result.slice
	}

	if result, err := msInterleave(parts[0], parts[1]); err != nil {
		resultCh <- &msWork{err: err}
	} else {
		resultCh <- &msWork{result, work.side, nil}
	}
}

// Merge implements a generic, concurrent sort utilizing the merge sort
// algorithm. The input slice is left unmodified and a sorted version is
// returned.
func Merge(slice gil.Slice) (gil.Slice, error) {
	workCh, resultCh := make(chan sortWork), make(chan sortWork)

	go ms(workCh, resultCh)
	workCh <- &msWork{slice: slice}

	result, ok := (<-resultCh).(*msWork)
	if !ok {
		return nil, gil.TypeAssertionError{}
	}

	return result.slice, result.err
}
