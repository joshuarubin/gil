package gil

import "fmt"

// RangeError

type RangeError struct {
	Interface
}

func (e RangeError) Error() string {
	return e.String() + " not in range"
}

// NotFoundError

type NotFoundError struct {
	Interface
}

func (e NotFoundError) Error() string {
	return e.String() + " not found"
}

// TypeMismatchError

type TypeMismatchError struct {
	lhs, rhs Interface
}

func (e TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch %s(%T):%s(%T)", e.lhs.String(), e.lhs, e.rhs.String(), e.rhs)
}

type EmptyError string

func (e EmptyError) Error() string {
	return string(e)
}
