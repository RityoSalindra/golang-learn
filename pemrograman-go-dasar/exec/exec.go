package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var ls, _ = exec.Command("cmd", "/C", "ls").Output()
	fmt.Printf(" -> ls\n%s\n", string(ls))

	var pwd, _ = exec.Command("cmd", "/C", "pwd").Output()
	fmt.Printf(" -> pwd\n%s\n", string(pwd))

	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf(" -> git config user.name\n%s\n", string(output3))

}
