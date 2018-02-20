package main

import (
	"flag"
	"fmt"
	"os"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
	"git.hduhelp.com/hduhelper/lecture/src/backend/routers"
)

var confpath = flag.String("confile", "conf/app.toml", "配置路径 默认 conf/app.toml")

func main() {
	flag.Parse()
	conf, err := conf.LoadConfig(*confpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := routers.SetupRouters()
	r.Run(conf.ListenAddr)
}
