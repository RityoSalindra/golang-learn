package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var output []byte
	var err error

	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", "ls").Output()
	} else {
		output, err = exec.Command("bash", "-c", "ls").Output() //karena pada windows , executable untuk ls dan beberapa command tidak ada sehingga menimbulkan error
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(output))
}
