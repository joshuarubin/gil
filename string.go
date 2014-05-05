package gil

// String is an string implementing Interface
type String string

// Less returns true if i String is smaller than iface Interface
func (i String) Less(iface Interface) (bool, error) {
	rhs, ok := iface.(String)
	if !ok {
		return false, TypeMismatchError{i, rhs}
	}

	return i < rhs, nil
}

// CopyToStringSlice takes a slice of ints and returns a slice of Interfaces.
func CopyToStringSlice(data []string) []Interface {
	ret := make([]Interface, len(data))
	for i, val := range data {
		ret[i] = String(val)
	}
	return ret
}
