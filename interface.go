package gil

type Interface interface {
	String() string // implement fmt.Stringer
	Less(iface Interface) (bool, error)
}
