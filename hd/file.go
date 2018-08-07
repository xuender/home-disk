package hd

import (
	"os"
	"strings"
	"time"

	"gopkg.in/h2non/filetype.v1"
)

type File struct {
	Id        string    `json:"id"`   // 主键
	Name      string    `json:"name"` // 文件名
	Ct        time.Time `json:"ct"`   // 创建时间
	Type      string    `json:"type"` // 类型
	Sub       string    `json:"sub"`  // 子类型
	Thumbnail []byte    `json:"-"`    // 缩略图
}

func NewFile(file string, size int) (*File, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	r := new(File)
	r.Name = fi.Name()
	r.Ct = fi.ModTime()
	t, err := filetype.MatchFile(file)
	if err == nil {
		r.Type = t.MIME.Type
		r.Sub = t.MIME.Subtype
		if r.Type == "image" {
			// 读取exif
			exif, err := NewExif(file)
			if err == nil {
				// 照片创建时间
				r.Ct = exif.DateTime
				// TODO 未来增加经纬度等
			}
			// 缩略图
			r.Thumbnail, _ = thumbnail(file, size, size)
		}
	}
	return r, nil
}

// 目录
func (r *File) Path(path string) string {
	return strings.Join([]string{
		path,
		r.Ct.Format("2006"),
		r.Ct.Format("01"),
		r.Ct.Format("02"),
	}, string(os.PathSeparator))
}

// 全称
func (r *File) FullName(path string) string {
	return strings.Join([]string{
		path,
		r.Ct.Format("2006"),
		r.Ct.Format("01"),
		r.Ct.Format("02"),
		r.Name,
	}, string(os.PathSeparator))
}

func (f *File) Day() string {
	return f.Ct.Format("2006-01-02")
}
