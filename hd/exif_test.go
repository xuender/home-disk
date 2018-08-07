package hd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExif(t *testing.T) {
	Convey("Exif", t, func() {
		Convey("i.jpg", func() {
			r, err := NewExif("i.jpg")
			So(r, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(r.DateTime.Format(time_format), ShouldEqual, "2018:07:22 09:26:37")
		})
	})
}
