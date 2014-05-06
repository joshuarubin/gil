package search

import (
	"sort"

	"github.com/joshuarubin/gil"
)

// BinarySearchSlice implements a generic binary search. The index of the matching value
// is returned. If no match is found, NotFoundError is returned.
func BinarySearchSlice(slice []gil.Interface, val gil.Interface) (int, error) {
	l := len(slice)

	pos := sort.Search(l, func(i int) bool {
		test := slice[i]
		return (val == test) || val.Less(slice[i])
	})

	if pos < l && slice[pos] == val {
		return pos, nil
	}

	return pos, gil.NotFoundError{Interface: val}
}
