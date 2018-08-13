package hd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/*
func TestSize(t *testing.T) {
	Convey("size", t, func() {
		w, h := size(100, 200, 10, 20)
		So(w, ShouldEqual, 10)
		So(h, ShouldEqual, 20)
		w, h = size(10, 20, 100, 200)
		So(w, ShouldEqual, 10)
		So(h, ShouldEqual, 20)
		w, h = size(100, 200, 20, 20)
		So(w, ShouldEqual, 20)
		So(h, ShouldEqual, 40)
		w, h = size(100, 200, 30, 10)
		So(w, ShouldEqual, 30)
		So(h, ShouldEqual, 60)
		w, h = size(10, 20, 100, 200)
		So(w, ShouldEqual, 10)
		So(h, ShouldEqual, 20)
	})
}
func TestClip(t *testing.T) {
	Convey("clip", t, func() {
		Convey("ok", func() {
			x0, y0, x1, y1 := clip(200, 200, 100, 200)
			So(y0, ShouldEqual, 0)
			So(y1, ShouldEqual, 200)
			So(x0, ShouldEqual, 50)
			So(x1, ShouldEqual, 150)
		})
		Convey("error", func() {
			x0, y0, x1, y1 := clip(20, 20, 100, 200)
			So(y0, ShouldEqual, 0)
			So(y1, ShouldEqual, 20)
			So(x0, ShouldEqual, 0)
			So(x1, ShouldEqual, 20)
		})
	})
}
*/
func TestThumbnail(t *testing.T) {
	Convey("thumbnail", t, func() {
		Convey("ok", func() {
			bs, err := thumbnail("i.jpg", 200, 200)
			So(err, ShouldBeNil)
			So(bs, ShouldNotBeNil)
		})
	})
}
