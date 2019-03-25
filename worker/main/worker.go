package main

import (
	"flag"
	"fmt"
	"github.com/coderJianXun/go_crontab/worker"
	"runtime"
	"time"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// worker-config ./master.josn
	// worker -h
	flag.StringVar(&confFile, "config", "./worker.json", "指定worker.json")
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
	if err = worker.InitConfig(confFile); err != nil {
		goto ERR
	}

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}
	return

ERR:
	fmt.Println(err)
}
