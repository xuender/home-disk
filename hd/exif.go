package hd

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	exif "github.com/m0rcq/go-exif"
)

const time_format = "2006:01:02 15:04:05"

type Exif struct {
	DateTime time.Time // 创建时间
	Make     string    // 相机
	Software string    // 软件
	Model    string    // 模式
	// GPSLatitude
	// GPSLongitudeRef
	// GPSLongitude
	// GPSAltitudeRef
}

func NewExif(file string) (ret *Exif, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	ret = new(Exif)
	exif := &exif.ExifData{}
	_, err = exif.GetExifData(f)
	if err != nil {
		return
	}
	exif.ProcessExifStream(f)
	ret.Make = getValue(exif, "Make")
	ret.Software = getValue(exif, "Software")
	ret.Model = getValue(exif, "Model")
	t, err := time.Parse(time_format, getValue(exif, "DateTime"))
	if err == nil {
		ret.DateTime = t
	} else {
		ret.DateTime = time.Now()
	}
	return
}
func getValue(exif *exif.ExifData, key string) string {
	for _, ifds := range exif.IfdData {
		for _, v := range ifds {
			if v.TagDesc != key {
				continue
			}
			lval, ok := v.Values.([]interface{})
			var values string
			if ok && len(lval) > 0 {
				switch val := lval[0].(type) {
				case string:
					values = fmt.Sprintf("%s", val)
				case byte:
					values = fmt.Sprintf("%#x", val)
				case []uint8:
					var lstr []string
					for _, v := range lval {
						lstr = append(lstr, fmt.Sprintf("%#x", v))
					}
					values = strings.Join(lstr, ", ")
				case int16:
					values = fmt.Sprintf("%d", val)
				case int32:
					values = fmt.Sprintf("%d", val)
				case int64:
					values = fmt.Sprintf("%d", val)
				case uint16:
					values = fmt.Sprintf("%d", val)
				case uint32:
					values = fmt.Sprintf("%d", val)
				case uint64:
					values = fmt.Sprintf("%d", val)
				case *big.Rat:
					values = fmt.Sprintf("%s", val.RatString())
				default:
					values = fmt.Sprintf("%v", lval)
				}
			}
			return values
		}
	}
	return ""
}
