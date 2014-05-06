package sort

import "github.com/joshuarubin/gil"

type sortWork interface {
	Slice() gil.Slice
	SetSlice(slice gil.Slice)
}
