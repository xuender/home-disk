package hd

import (
	"bytes"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/bakape/thumbnailer"
)

// 缩略图
func thumbnail(file string, width, height int) ([]byte, error) {
	// 读取原图
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	o := thumbnailer.Options{
		ThumbDims: thumbnailer.Dims{
			Width:  uint(width),
			Height: uint(height),
		},
	}
	_, thumb, err := thumbnailer.Process(f, o)
	if err != nil {
		return nil, err
	}
	return thumb.Data, nil
}

// 播放标记
func play(ts, ps []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewBuffer(ts))
	if err != nil {
		return nil, err
	}
	wmb_img, err := png.Decode(bytes.NewBuffer(ps))
	if err != nil {
		return nil, err
	}
	//把水印写在右下角，并向0坐标偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-10, img.Bounds().Dy()-wmb_img.Bounds().Dy()-10)
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, m, &jpeg.Options{100})
	return buf.Bytes(), nil
}

/*
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
	if ow > nw {
		x0 = (ow - nw) / 2
		x1 = ow - x0
	} else {
		x0 = 0
		x1 = ow
	}
	if oh > nh {
		y0 = (oh - nh) / 2
		y1 = oh - y0
	} else {
		y0 = 0
		y1 = oh
	}
	return
}
*/
