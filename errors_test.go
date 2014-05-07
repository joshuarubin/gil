package gil

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {
	Convey("Errors should work properly", t, func() {
		So(NotFoundError{Int(0)}.Error(), ShouldEqual, "not found")
		So(TypeAssertionError{}.Error(), ShouldEqual, "type assertion error")
		So(ArgumentError("some error text").Error(), ShouldEqual, "some error text")
	})
}
