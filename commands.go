package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var CommandList = map[string]string{
	"CWD": "cd",
	"LS":  "ls",
}

func (c Client) RunCommand(cmd string) string {
	words := strings.Fields(cmd)
	if val, ok := CommandList[words[0]]; ok {
		fmt.Println("homedir: ", c.homeDir)
		fmt.Println("got: ", words)
		// TODO; Check if len > 1
		shcmd := val
		if len(words) > 1 {
			shcmd += strings.Join(words[1:], " ")
		}
		// Run command...
		fmt.Print("Trying to run '" + shcmd + "'")
		out, err := exec.Command(shcmd).Output()
		if err != nil {
			return "fail"
		}
		fmt.Println("out: " + string(out))
		return "run " + val + " and got " + string(out)
	}
	return "No such command!"
}
