package hd

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNewResource(t *testing.T) {
	Convey("NewResource", t, func() {
		Convey("error file", func() {
			r, err := NewResource("/aaaaaa")
			So(err, ShouldNotBeNil)
			So(r, ShouldBeNil)
		})
		Convey("i.jpg", func() {
			r, err := NewResource("i.jpg")
			So(r, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(r.Name, ShouldEqual, "i.jpg")
			So(r.Mtype, ShouldEqual, "image")
			So(r.Msub, ShouldEqual, "jpeg")
			t, _ := time.Parse(time_format, "2018:07:22 09:26:37")
			So(r.Ct, ShouldEqual, t)
			So(r.Path("/tmp"), ShouldEqual, "/tmp/2018/07/22")
			So(r.FullName("/tmp"), ShouldEqual, "/tmp/2018/07/22/i.jpg")
		})
	})
}
