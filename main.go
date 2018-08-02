package main

import (
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/xuender/goutils"

	"./hd"
)

func main() {
	app := cli.NewApp()
	app.Name = "home-disk"
	app.Usage = "家庭数据盘"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "browser,b",
			Usage: "启动禁止打开浏览器",
		},
		cli.StringFlag{
			Name:  "port,p",
			Value: "6181",
			Usage: "访问端口",
		},
		cli.StringFlag{
			Name:  "temp,t",
			Value: "temp",
			Usage: "临时文件目录",
		},
		cli.StringFlag{
			Name:  "data,d",
			Value: "data",
			Usage: "文件存储目录",
		},
		cli.StringFlag{
			Name:  "storage,s",
			Value: "storage",
			Usage: "数据库目录",
		},
	}
	app.Action = func(c *cli.Context) error {
		var port = c.String("port")
		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
		// 打开浏览器
		if !c.Bool("browser") {
			url, err := hd.GetUrl(port)
			if err == nil {
				goutils.Open(url)
			}
		}
		web := hd.Web{
			Port: port,
			Temp: c.String("temp"),
			Data: c.String("data"),
			Storage: c.String("storage"),
		}
		web.Run()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
