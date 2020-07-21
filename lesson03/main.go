package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	var (
		cmd        *exec.Cmd
		ctx        context.Context
		CancelFunc context.CancelFunc
		err        error
		output     []byte
	)

	ctx, CancelFunc = context.WithCancel(context.TODO())
	go func() {
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 3; ls -l")
		output, err = cmd.CombinedOutput()

		fmt.Println("协程输出：" + string(output))

	}()

	time.Sleep(6 * time.Second)
	println("主程结束")
	CancelFunc()

}
