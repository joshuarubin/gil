package gil

// Interface is the interface used for storing data generically within gil
type Interface interface {
	Less(iface Interface) (bool, error)
}

// Queue is a generic queue interface implemented by some types in gil/containers
type Queue interface {
	Len() int
	Push(item interface{}) Queue
	Pop() interface{}
	Peek() interface{}
}
