package main

import (
	"flag"

	"../../api/demo/v1"
	"../../conf"
	"../../internal/service"
)

func main() {

	// 初始化配置文件
	var conf_path string
	flag.StringVar(&conf_path, "conf", "conf/app.ini", "program config file")
	flag.Parse()
	conf.SetUp(conf_path)

	// 初始化数据库
	service.Setup()

	//
	demo.DemoInit()
}
