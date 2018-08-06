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
	"sort"
)

type Web struct {
	Port string      // 端口号
	Temp string      // 临时文件目录
	Data string      // 保存数据目录
	Db   string      // 数据库目录
	Size int         // 缩略图尺寸
	db   *leveldb.DB // 数据库
	days Days        // 文件日期列表
}

// 文件日期列表主键
var DAYS_KEY = []byte("days")

func (w *Web) Init(reset bool) error {
	if reset {
		// 删除数据库
		log.Printf("删除数据库 %s\n", w.Db)
		os.RemoveAll(w.Db)
	}
	db, err := leveldb.OpenFile(w.Db, nil)
	// 数据库链接
	w.db = db
	if err != nil {
		return err
	}
	// days 读取
	data, err := db.Get(DAYS_KEY, nil)
	if err == nil {
		goutils.Decode(data, &w.days)
	} else {
		w.days = Days{}
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
			key, err := w.createKey(file)
			if err != nil {
				return err
			}
			_, err = w.save(file, key)
			return err
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
	e.GET("/days", w.getDays)
	e.GET("/days/:day", w.getDay)
	// 缩略图
	e.GET("/t/:id", w.thumbnail)
	// 文件信息
	e.GET("/file/:id", w.getFile)
	// 下载
	e.GET("/down/:id", w.download)
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
	key, err := w.createKey(file)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// 查找DB
	if w.isOld(key) {
		os.Remove(file)
		return c.String(http.StatusOK, "重复")
	}
	f, err := w.save(file, key)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, f.Id)
}

func (w *Web) save(file string, key []byte) (*File, error) {
	// 文件信息
	f, err := NewFile(file, w.Size)
	f.Id = hex.EncodeToString(key)
	if err != nil {
		return f, err
	}
	// 保存文件
	mkdir(f.Path(w.Data))
	if file != f.FullName(w.Data) {
		os.Rename(file, f.FullName(w.Data))
	}
	// 保存文件列表
	day := f.Day()
	w.addList(day, key)
	// 文件信息保存
	bs, _ := goutils.Encode(f)
	w.db.Put(key, bs, nil)
	// 追加日期列表
	if w.days.Add(day) {
		bs, _ := goutils.Encode(w.days)
		w.db.Put(DAYS_KEY, bs, nil)
	}
	return f, nil
}

func (w *Web) addList(day string, key []byte) {
	keys := [][]byte{}
	k := []byte(day)
	data, err := w.db.Get(k, nil)
	if err == nil {
		goutils.Decode(data, &keys)
	}
	keys = append(keys, key)
	bs, _ := goutils.Encode(keys)
	w.db.Put(k, bs, nil)
}

func (w *Web) isOld(key []byte) bool {
	// 查找DB
	data, err := w.db.Get(key, nil)
	if err == nil {
		// 旧文件
		nr := File{}
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

func (w *Web) createKey(file string) ([]byte, error) {
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
	r := File{}
	goutils.Decode(data, &r)
	return c.Blob(http.StatusOK, "image/jpeg", r.Thumbnail)
}

// 缩略图
func (w *Web) getFile(c echo.Context) error {
	id := c.Param("id")
	key, err := hex.DecodeString(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "文件ID错误: "+id)
	}
	data, err := w.db.Get(key, nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, "缩略图数据错误: "+id)
	}
	r := File{}
	goutils.Decode(data, &r)
	return c.JSON(http.StatusOK, r)
}

// 下载文件
func (w *Web) download(c echo.Context) error {
	id := c.Param("id")
	key, err := hex.DecodeString(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "下载ID错误: "+id)
	}
	data, err := w.db.Get(key, nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, "缩略图数据错误: "+id)
	}
	f := File{}
	goutils.Decode(data, &f)
	return c.File(f.FullName(w.Data))
}

func (w *Web) getDays(c echo.Context) error {
	return c.JSON(http.StatusOK, w.days)
}

// 获取日文件信息列表
func (w *Web) getDay(c echo.Context) error {
	day := c.Param("day")
	keys := [][]byte{}
	k := []byte(day)
	data, err := w.db.Get(k, nil)
	if err == nil {
		goutils.Decode(data, &keys)
		ret := []File{}
		for _, key := range keys {
			data, err := w.db.Get(key, nil)
			if err == nil {
				f := File{}
				goutils.Decode(data, &f)
				ret = append(ret, f)
			}
		}
		sort.SliceStable(ret, func(i, j int) bool { return ret[i].Ct.Unix() < ret[j].Ct.Unix() })
		return c.JSON(http.StatusOK, ret)
	}
	return c.JSON(http.StatusOK, []string{})
}
