package gil

// Interface is the interface used for storing data generically within gil
type Interface interface {
	String() string // implement fmt.Stringer
	Less(iface Interface) (bool, error)
}
