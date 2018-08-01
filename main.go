package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"rsc.io/qr"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xuender/goutils"
)

const PORT = ":1323"

func main() {
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
	e.GET("/qr", qrcode)
	e.GET("/home", hello)
	e.GET("/info", ip)
	e.POST("/up", upload)

	// 打开浏览器
	url, err := getUrl()
	if err == nil {
		goutils.Open(url)
	}

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}

func upload(c echo.Context) error {
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

	_, err = os.Stat("data")
	if os.IsNotExist(err) {
		os.Mkdir("data", 0777)
	}
	// 目的
	dst, err := os.Create("data" + string(os.PathSeparator) + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 复制
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.String(http.StatusOK, "File is uploaded")
}

// 测试
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// IP
func ip(c echo.Context) error {
	ip, err := getIp()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, ip)
}

// 获取IP地址
func getIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && !strings.HasPrefix(ipnet.IP.String(), "172") {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("未找到IP")
}

// QR码
func qrcode(c echo.Context) error {
	url, err := getUrl()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	code, err := qr.Encode(url, qr.Q)
	if err != nil {
		return c.String(http.StatusInternalServerError, "QR码生成错误")
	}
	return c.Blob(http.StatusOK, "image/png", code.PNG())
}

func getUrl() (string, error) {
	ip, err := getIp()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://%s%s", ip, PORT), nil
}
