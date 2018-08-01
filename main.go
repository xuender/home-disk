package main

import (
	"log"
	"os"

	"./hd"
	"github.com/urfave/cli"
	"github.com/xuender/goutils"
)

func main() {
	app := cli.NewApp()
	app.Name = "home-disk"
	app.Usage = "家庭数据盘"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "browser,b",
			Usage: "启动禁止打开浏览器",
		},
	}
	app.Action = func(c *cli.Context) error {
		var port = ":6181"
		// 打开浏览器
		if !c.Bool("browser") {
			url, err := hd.GetUrl(port)
			if err == nil {
				goutils.Open(url)
			}
		}
		web := hd.Web{Port: port, Temp: "tmp", Data: "data"}
		web.Run()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
