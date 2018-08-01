package hd

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xuender/goutils"
	"io"
	"log"
	"net/http"
	"os"
	"rsc.io/qr"
)

type Web struct {
	Port string // 端口号
	Temp string // 临时文件目录
	Data string // 保存数据目录
}

func (w *Web) Run() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.POST},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	e.Static("/", "www")
	// Routes
	e.GET("/qr", w.qrcode)
	e.POST("/up", w.upload)

	// Start server
	e.Logger.Fatal(e.Start(w.Port))
}
func (w *Web) upload(c echo.Context) error {
	c.Logger().Warn("文件上传")
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	mkdir(w.Data)
	mkdir(w.Temp)
	// 目的
	f := w.Temp + string(os.PathSeparator) + file.Filename
	dst, err := os.Create(f)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 复制
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	// TODO 保存FileId
	fid, err := goutils.NewFileId(f)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	log.Println(fid)
	r, err := NewResource(f)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	mkdir(r.Path(w.Data))
	os.Rename(f, r.FullName(w.Data))
	return c.String(http.StatusOK, "File is uploaded")
}

// QR码
func (w *Web) qrcode(c echo.Context) error {
	url, err := GetUrl(w.Port)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	code, err := qr.Encode(url, qr.Q)
	if err != nil {
		return c.String(http.StatusInternalServerError, "QR码生成错误")
	}
	return c.Blob(http.StatusOK, "image/png", code.PNG())
}
