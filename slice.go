package gil

import "sort"

type Slice []Interface

func (s Slice) Find(val Interface) int {
	l := len(s)

	return sort.Search(l, func(i int) bool {
		test := s[i]
		return (val == test) || val.Less(s[i])
	})
}

// required for sort.Interface
func (s Slice) Len() int {
	return len(s)
}

// required for sort.Interface
func (s Slice) Less(i, j int) bool {
	return s[i].Less(s[j])
}

// required for sort.Interface
func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort the slice
func (s Slice) Sort() {
	sort.Sort(s)
}
