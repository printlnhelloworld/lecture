package main

import (
	"flag"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
	"git.hduhelp.com/hduhelper/lecture/src/backend/routers"
)

var confpath = flag.String("confile", "conf/app.toml", "配置路径 默认 conf/app.toml")

func main() {
	flag.Parse()
	conf, err := conf.LoadConfig(*confpath)
	checkErr(err)

	err = model.InitDB(conf)
	checkErr(err)

	r := routers.SetupRouters()
	r.Run(conf.ListenAddr)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
