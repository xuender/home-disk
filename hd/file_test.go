package hd

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFile(t *testing.T) {
	Convey("NewFile", t, func() {
		Convey("error file", func() {
			r, err := NewFile("/aaaaaa", 20)
			So(err, ShouldNotBeNil)
			So(r, ShouldBeNil)
		})
		Convey("i.jpg", func() {
			r, err := NewFile("i.jpg", 20)
			So(r, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(r.Name, ShouldEqual, "i.jpg")
			So(r.Type, ShouldEqual, "image")
			So(r.Sub, ShouldEqual, "jpeg")
			t, _ := time.Parse(time_format, "2018:07:22 09:26:37")
			So(r.Ct, ShouldEqual, t)
			So(r.Path("/tmp"), ShouldEqual, "/tmp/2018/07/22")
			So(r.FullName("/tmp"), ShouldEqual, "/tmp/2018/07/22/i.jpg")
		})
	})
}
