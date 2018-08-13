package hd

import (
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/zhulik/go_mediainfo"
)

type Mediainfo struct {
	DateTime time.Time // 创建时间
	Duration int       // 时长, 单位毫秒
}

const utc_time_format = "UTC 2006-01-02 15:04:05"

func NewMediainfo(file string) (ret *Mediainfo, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	mi := mediainfo.NewMediaInfo()
	err = mi.OpenMemory(bytes)
	if err != nil {
		return
	}

	ret = new(Mediainfo)
	// 创建时间
	ed := mi.Get("Encoded_Date")
	if ed != "" {
		t, err := time.Parse(utc_time_format, ed)
		if err == nil {
			ret.DateTime = t
		}
	}
	// 时长
	d := mi.Get("Duration")
	if d != "" {
		ret.Duration, _ = strconv.Atoi(d)
	}
	return
}
