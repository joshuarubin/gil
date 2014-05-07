package gil

// Interface is the interface used for storing data generically within gil
type Interface interface {
	Less(iface Interface) bool
}

// Queue is a generic queue interface implemented by some types in gil/containers
type QueueInterface interface {
	Len() int
	Push(item Interface) QueueInterface
	Pop() Interface
	Peek() Interface
}
