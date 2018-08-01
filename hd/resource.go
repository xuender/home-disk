package hd

import (
	"github.com/xiam/exif"
	"gopkg.in/h2non/filetype.v1"
	"os"
	"strings"
	"time"
)

type Resource struct {
	Name  string    // 文件名
	Ct    time.Time // 创建时间
	Mtype string    // 类型
	Msub  string    // 子类型
}

const time_format = "2006:01:02 15:04:05"

func NewResource(file string) (*Resource, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	r := new(Resource)
	r.Name = fi.Name()
	r.Ct = fi.ModTime()
	t, err := filetype.MatchFile(file)
	if err == nil {
		r.Mtype = t.MIME.Type
		r.Msub = t.MIME.Subtype
		if r.Mtype == "image" {
			// 读取exif
			data, err := exif.Read(file)
			if err == nil {
				// 照片创建时间
				ct := data.Tags["Date and Time"]
				if len(ct) == 19 {
					t, err := time.Parse(time_format, ct)
					if err == nil {
						r.Ct = t
					}
				}
				// TODO 未来增加经纬度等
			}
		}
	}
	return r, nil
}

// 目录
func (r *Resource) Path(path string) string {
	return strings.Join([]string{
		path,
		r.Ct.Format("2006"),
		r.Ct.Format("01"),
		r.Ct.Format("02"),
	}, string(os.PathSeparator))
}

// 全称
func (r *Resource) FullName(path string) string {
	return strings.Join([]string{
		path,
		r.Ct.Format("2006"),
		r.Ct.Format("01"),
		r.Ct.Format("02"),
		r.Name,
	}, string(os.PathSeparator))
}
