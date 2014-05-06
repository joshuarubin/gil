package gil

import "sort"

// Interface is the interface used for storing data generically within gil
type Interface interface {
	Less(iface Interface) bool
}

type Slice []Interface

func (s Slice) Find(val Interface) (int, error) {
	l := len(s)

	pos := sort.Search(l, func(i int) bool {
		test := s[i]
		return (val == test) || val.Less(s[i])
	})

	if pos < l && s[pos] == val {
		return pos, nil
	}

	return pos, NotFoundError{Interface: val}
}

// Queue is a generic queue interface implemented by some types in gil/containers
type Queue interface {
	Len() int
	Push(item Interface) Queue
	Pop() Interface
	Peek() Interface
}
