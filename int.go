package gil

import "strconv"

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (lhs Int) Less(iface Interface) (bool, error) {
	rhs, ok := iface.(Int)
	if !ok {
		return false, TypeMismatchError{lhs, rhs}
	}

	return lhs < rhs, nil
}

func CopyToIntSlice(data []int) []Interface {
	ret := make([]Interface, len(data))
	for i, val := range data {
		ret[i] = Int(val)
	}
	return ret
}
