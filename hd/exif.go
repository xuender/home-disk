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
	err = processExifStream(f, exif)
	if err != nil {
		return
	}

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
func processExifStream(f *os.File, self *exif.ExifData) error {
	data, err := self.GetExifData(f)

	if err != nil {
		return err
	}

	var tiff exif.TiffInterface = &exif.TiffData{}

	tiff.SetByteOrder(string(data[0]))
	tiff.SetData(data)
	ifdList := tiff.GetIFDList()

	self.IfdData = make(map[string][]exif.IfdEntries, len(ifdList)+3) // provision for SubIFD, GPSInfo, IOPInfo

	for k, v := range ifdList {
		ifds := tiff.ProcessIFD(uint8(k), v, exif.EXIF_TAGS)
		self.IfdData[exif.IfdSeqMap[uint8(k)]] = ifds

		var subIfdOffset uint32
		var gpsInfoOffset uint32
		var iopOffset uint32

		for _, v := range ifds {
			if v.Tag == exif.EXIF_SUBIFD_OFFSET_TAG {
				values, ok := v.Values.([]interface{})
				if ok {
					val, ok := values[0].(uint32)
					if ok {
						subIfdOffset = uint32(val)
					}
				}
			} else if v.Tag == exif.EXIF_GPSINFO_TAG {
				values, ok := v.Values.([]interface{})
				if ok {
					val, ok := values[0].(uint32)
					if ok {
						gpsInfoOffset = uint32(val)
					}
				}
			}
		}

		if subIfdOffset != 0 {
			subIfd := tiff.ProcessIFD(2, subIfdOffset, exif.EXIF_TAGS)
			self.IfdData[exif.IfdSeqMap[2]] = subIfd
			for _, v := range subIfd { // try to find SubIFD Interoperability Tag
				if v.Tag == exif.EXIF_IOP_TAG {
					values, ok := v.Values.([]interface{})
					if ok {
						val, ok := values[0].(uint32)
						if ok {
							iopOffset = uint32(val)
						}
					}
				}
			}
		}

		if gpsInfoOffset != 0 {
			gpsInfo := tiff.ProcessIFD(3, gpsInfoOffset, exif.EXIF_GPSINFO_TAGS)
			self.IfdData[exif.IfdSeqMap[3]] = gpsInfo
		}

		if iopOffset != 0 {
			iopInfo := tiff.ProcessIFD(4, iopOffset, exif.EXIF_IOP_TAGS)
			self.IfdData[exif.IfdSeqMap[4]] = iopInfo
		}
	}
	return nil
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
