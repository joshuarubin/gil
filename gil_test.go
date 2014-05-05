package gil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	list := CopyToIntSlice([]int{0, 1, 2, 3, 4, 5})
	Convey("CopyToIntSlice should work", t, func() {
		for i := 0; i < len(list); i++ {
			So(list[i], ShouldEqual, i)
		}
	})
}
