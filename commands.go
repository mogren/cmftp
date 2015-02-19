package main

import (
	"fmt"
)

var CommandList = map[string]string{
	"CWD": "cd",
	"LS":  "ls",
}

func RunCommand(cmd string) string {
	if val, ok := CommandList[cmd]; ok {
		fmt.Println("got " + cmd)
		return "run " + val
	}
	return "No such command!"
}
