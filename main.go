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
		cli.StringFlag{
			Name:  "data,d",
			Value: "data",
			Usage: "文件存储目录",
		},
		cli.StringFlag{
			Name:  "db,b",
			Value: "db",
			Usage: "数据库目录",
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
		cli.IntFlag{
			Name:  "size,s",
			Value: 200,
			Usage: "缩略图尺寸",
		},
		cli.BoolFlag{
			Name:  "no-open,n",
			Usage: "启动不打开浏览器",
		},
		cli.BoolFlag{
			Name:  "db-reset,r",
			Usage: "数据库重置",
		},
	}
	app.Action = func(c *cli.Context) error {
		var port = c.String("p")
		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
		web := hd.Web{
			Port: port,
			Temp: c.String("t"),
			Data: c.String("d"),
			Db:   c.String("b"),
			Size: c.Int("s"),
		}
		// 初始化
		err := web.Init(c.Bool("r"))
		if err != nil {
			return err
		}
		// 打开浏览器
		if !c.Bool("n") {
			url, err := hd.GetUrl(port)
			if err == nil {
				goutils.Open(url)
			}
		}
		// 运行
		web.Run()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
