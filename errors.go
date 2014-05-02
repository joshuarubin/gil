package gil

import "fmt"

// RangeError is returned when a search value lies outside the range of
// the set of possibilities
type RangeError struct {
	Interface
}

func (e RangeError) Error() string {
	return e.String() + " not in range"
}

// NotFoundError is returned when a search value is not found within a set
type NotFoundError struct {
	Interface
}

func (e NotFoundError) Error() string {
	return e.String() + " not found"
}

// TypeAssertionError is returned when an internal type assertion failed.
// It indicates a library problem and should be reported.
// It should not be seen under normal circumstances
type TypeAssertionError struct{}

func (e TypeAssertionError) Error() string {
	return "type assertion error"
}

// TypeMismatchError is returned when trying to compare two values with
// different types
type TypeMismatchError struct {
	lhs, rhs Interface
}

func (e TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch %s(%T):%s(%T)", e.lhs.String(), e.lhs, e.rhs.String(), e.rhs)
}

// ArgumentError is returned when an invalid argument or number of arguments
// is passed to a function, often variadic.
type ArgumentError string

func (e ArgumentError) Error() string {
	return string(e)
}
