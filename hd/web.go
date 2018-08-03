package hd

import (
	"encoding/hex"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/xuender/goutils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"rsc.io/qr"
)

type Web struct {
	Port string      // 端口号
	Temp string      // 临时文件目录
	Data string      // 保存数据目录
	Db   string      // 数据库目录
	Size int         // 缩略图尺寸
	db   *leveldb.DB // 数据库
}

func (w *Web) Init(reset bool) error {
	if reset {
		// 删除数据库
		log.Printf("删除数据库 %s\n", w.Db)
		os.RemoveAll(w.Db)
	}
	db, err := leveldb.OpenFile(w.Db, nil)
	w.db = db
	if err != nil {
		return err
	}
	if reset {
		log.Println("数据库重置")
		return filepath.Walk(w.Data, func(file string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			key, err := w.getKey(file)
			if err != nil {
				return err
			}
			return w.save(file, key)
		})
	}
	return nil
}

func (w *Web) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.POST},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	// 静态资源
	e.Static("/", "www")
	// 二维码
	e.GET("/qr", w.qrcode)
	// 缩略图
	e.GET("/t/:id", w.thumbnail)
	// 文件上传
	e.POST("/up", w.upload)
	// 启动服务
	e.Logger.Fatal(e.Start(w.Port))
}

// 上传文件
func (w *Web) upload(c echo.Context) error {
	file, err := w.saveTemp(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// 生成文件标识
	key, err := w.getKey(file)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// 查找DB
	if w.isOld(key) {
		os.Remove(file)
		return c.String(http.StatusOK, "重复")
	}
	err = w.save(file, key)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "OK")
}

func (w *Web) save(file string, key []byte) error {
	log.Println(hex.EncodeToString(key))
	r, err := NewResource(file, w.Size)
	if err != nil {
		return err
	}
	mkdir(r.Path(w.Data))
	if file != r.FullName(w.Data) {
		os.Rename(file, r.FullName(w.Data))
	}
	bs, _ := goutils.Encode(r)
	w.db.Put(key, bs, nil)
	return nil
}

func (w *Web) isOld(key []byte) bool {
	// 查找DB
	data, err := w.db.Get(key, nil)
	if err == nil {
		// 旧文件
		nr := Resource{}
		goutils.Decode(data, &nr)
		_, err := os.Stat(nr.FullName(w.Data))
		return err == nil
	}
	return false
}

func (w *Web) saveTemp(c echo.Context) (string, error) {
	// 来源
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// 目的
	mkdir(w.Temp)
	f := w.Temp + string(os.PathSeparator) + file.Filename
	dst, err := os.Create(f)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// 复制
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return f, nil
}

func (w *Web) getKey(file string) ([]byte, error) {
	fid, err := goutils.NewFileId(file)
	if err != nil {
		return nil, err
	}
	return goutils.PrefixBytes([]byte("R-"), fid.Id()), nil
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

// 缩略图
func (w *Web) thumbnail(c echo.Context) error {
	id := c.Param("id")
	key, err := hex.DecodeString(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "缩略图ID错误: "+id)
	}
	data, err := w.db.Get(key, nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, "缩略图数据错误: "+id)
	}
	r := Resource{}
	goutils.Decode(data, &r)
	return c.Blob(http.StatusOK, "image/jpeg", r.Thumbnail)
}
