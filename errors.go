package gil

// NotFoundError is returned when a search value is not found within a set
type NotFoundError struct {
	Interface
}

func (NotFoundError) Error() string {
	return "not found"
}

// TypeAssertionError is returned when an internal type assertion failed.
// It indicates a library problem and should be reported.
// It should not be seen under normal circumstances
type TypeAssertionError struct{}

func (e TypeAssertionError) Error() string {
	return "type assertion error"
}

// ArgumentError is returned when an invalid argument or number of arguments
// is passed to a function, often variadic.
type ArgumentError string

func (e ArgumentError) Error() string {
	return string(e)
}
