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
	var (
		cmd        *exec.Cmd
		ctx        context.Context
		CancelFunc context.CancelFunc
		err        error
		output     []byte
		resultChan chan *result
	)
	resultChan = make(chan *result, 1000)

	ctx, CancelFunc = context.WithCancel(context.TODO())
	go func() {
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 3; ls -l")
		//执行任务捕获输出
		output, err = cmd.CombinedOutput()

		//fmt.Println("协程输出：" + string(output))
		//把任务输出结果传递main主程序
		resultChan <- &result{
			err:    err,
			output: output,
		}

	}()

	time.Sleep(6 * time.Second)
	println("主程结束")
	//取消上下文
	CancelFunc()
	if res := <-resultChan;res.err == nil{
		fmt.Println( string(res.output))
	}


}
