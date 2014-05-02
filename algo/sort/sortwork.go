package sort

import "github.com/joshuarubin/gil"

type sortWork interface {
	List() []gil.Interface
	SetList(list []gil.Interface)
}
