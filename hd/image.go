package hd

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"os"
)

// 缩略图
func thumbnail(pic string, width, height int) ([]byte, error) {
	// 读取原图
	f, err := os.Open(pic)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// 解码图片
	origin, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	// 缩放
	w, h := size(origin.Bounds().Max.X, origin.Bounds().Max.Y, width, height)
	canvas := resize.Thumbnail(uint(w), uint(h), origin, resize.Lanczos3)
	// 裁剪
	img := canvas.(*image.YCbCr)
	subImg := img.SubImage(image.Rect(clip(w, h, width, height))).(*image.YCbCr)
	// 输出JPG
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, subImg, &jpeg.Options{100})
	return buf.Bytes(), err
}

func size(ow, oh, nw, nh int) (int, int) {
	w := nw
	if ow < w {
		w = ow
	}
	h := nh
	if oh < h {
		h = oh
	}
	b := ow * 100 / oh
	if w < h*b/100 {
		w = h * b / 100
	}
	if h < w*100/b {
		h = w * 100 / b
	}
	return w, h
}

func clip(ow, oh, nw, nh int) (x0, y0, x1, y1 int) {
	if ow == nw {
		x0 = 0
		x1 = ow
	}
	if oh == nh {
		y0 = 0
		y1 = oh
	}
	if ow > nw {
		x0 = (ow - nw) / 2
		x1 = ow - x0
	}
	if oh > nh {
		y0 = (oh - nh) / 2
		y1 = oh - y0
	}
	return
}
