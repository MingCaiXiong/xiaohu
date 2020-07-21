package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
		output []byte
	)
	cmd = exec.Command("/bin/bash", "-c", "sleep 1; ls -l")
	output,err = cmd.CombinedOutput();
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))

}
