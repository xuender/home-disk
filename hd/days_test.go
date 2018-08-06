package hd

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDays(t *testing.T) {
	Convey("days", t, func() {
		days := Days{}
		So(len(days), ShouldEqual, 0)
		days.Add("1")
		So(len(days), ShouldEqual, 1)
		days.Add("1")
		So(len(days), ShouldEqual, 1)
		days.Add("2")
		So(len(days), ShouldEqual, 2)
		So(days[0], ShouldEqual, "2")
		So(days[1], ShouldEqual, "1")
	})
}
