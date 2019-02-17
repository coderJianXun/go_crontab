package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	// 执行一个cmd,让他在一个协程里去执行,让他执行2秒
	// 1秒的时候,杀死cmd
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	// 创建一个结果队列
	resultChan = make(chan *result, 1000)

	ctx, cancelFunc = context.WithCancel(context.TODO()) // ctx: 感知通道是否关闭, cancelFunc: 关闭通道
	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo hello")

		// 执行任务, 捕获输出
		output, err = cmd.CombinedOutput()

		// 把任务输出结果,传给main协程
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)

	cancelFunc()

	// 在main协程里,等待子协程的退出,并打印任务执行结果
	res = <-resultChan

	// 打印任务执行结果
	fmt.Println(res.err, string(res.output))
}
