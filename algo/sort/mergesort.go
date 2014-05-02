package sort

import "github.com/joshuarubin/gil"

type msData struct {
	pos  int
	list []gil.Interface
}

func (m *msData) Peek() gil.Interface {
	return m.list[m.pos]
}

func (m *msData) Pop() gil.Interface {
	ret := m.Peek()
	m.pos++
	return ret
}

func msInterleave(parts ...[]gil.Interface) ([]gil.Interface, error) {
	const NUM = 2

	if len(parts) != NUM {
		return nil, gil.ArgumentError("wrong number of parts")
	}

	data := [NUM]*msData{}
	tlen := 0
	for side, part := range parts {
		data[side] = &msData{list: part}
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
			var min *msData
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

type msWork struct {
	list []gil.Interface
	side int
	err  error
}

func (w *msWork) List() []gil.Interface {
	return w.list
}

func (w *msWork) SetList(list []gil.Interface) {
	w.list = list
}

func isShortList(work sortWork, copyList bool) bool {
	l := len(work.List())

	if l >= 2 {
		return false
	}

	if copyList {
		ret := make([]gil.Interface, l)
		copy(ret, work.List())
		work.SetList(ret)
	}

	return true
}

func ms(workCh <-chan sortWork, resultCh chan<- sortWork) {
	work, ok := (<-workCh).(*msWork)
	if !ok {
		resultCh <- &msWork{err: gil.TypeAssertionError{}}
		return
	}

	if isShortList(work, true) {
		resultCh <- work
		return
	}

	half := len(work.list) / 2
	parts := [][]gil.Interface{
		work.list[:half],
		work.list[half:],
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

		parts[result.side] = result.list
	}

	if result, err := msInterleave(parts[0], parts[1]); err != nil {
		resultCh <- &msWork{err: err}
	} else {
		resultCh <- &msWork{result, work.side, nil}
	}
}

// Merge implements a generic, concurrent sort utilizing the merge sort
// algorithm. The input list is left unmodified and a sorted version is
// returned.
func Merge(list []gil.Interface) ([]gil.Interface, error) {
	workCh, resultCh := make(chan sortWork), make(chan sortWork)

	go ms(workCh, resultCh)
	workCh <- &msWork{list: list}

	result, ok := (<-resultCh).(*msWork)
	if !ok {
		return nil, gil.TypeAssertionError{}
	}

	return result.list, result.err
}
