package main

import (
	"flag"
	"fmt"
	"github.com/coderJianXun/go_crontab/master"
	"runtime"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// master-config ./master.josn
	// master -h
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

// 初始化线程数量
func initEnv() {
	// 最大线程
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	// 初始化命令行参数
	initArgs()

	// 初始化线程
	initEnv()

	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	// 启动 Api Http 服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	return

ERR:
	fmt.Println(err)
}
