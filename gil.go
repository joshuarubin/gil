package gil

import (
	"fmt"
	"strconv"
)

type Interface interface {
	String() string // implement fmt.Stringer
	Less(iface Interface) (bool, error)
}

// ERRORS

type RangeError struct {
	Interface
}

func (e RangeError) Error() string {
	return e.String() + " not in range"
}

type NotFound struct {
	Interface
}

func (e NotFound) Error() string {
	return e.String() + " not found"
}

type TypeMismatch struct {
	lhs, rhs Interface
}

func (e TypeMismatch) Error() string {
	return fmt.Sprintf("type mismatch %s(%T):%s(%T)", e.lhs.String(), e.lhs, e.rhs.String(), e.rhs)
}

// TYPES

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (lhs Int) Less(iface Interface) (bool, error) {
	rhs, ok := iface.(Int)
	if !ok {
		return false, TypeMismatch{lhs, rhs}
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
