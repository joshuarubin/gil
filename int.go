package gil

// Int is an int implementing Interface
type Int int

// Less returns true if i Int is smaller than iface Interface
func (i Int) Less(iface Interface) (bool, error) {
	rhs, ok := iface.(Int)
	if !ok {
		return false, TypeMismatchError{i, rhs}
	}

	return i < rhs, nil
}

// CopyToIntSlice takes a slice of ints and returns a slice of Interfaces.
func CopyToIntSlice(data []int) []Interface {
	ret := make([]Interface, len(data))
	for i, val := range data {
		ret[i] = Int(val)
	}
	return ret
}
