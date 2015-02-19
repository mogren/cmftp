package main

import (
	"fmt"
	"strings"
)

var CommandList = map[string]string{
	"CWD": "cd",
	"LS":  "ls",
}

func (c Client) RunCommand(cmd string) string {
	words := strings.Fields(cmd)
	if val, ok := CommandList[words[0]]; ok {
		fmt.Println("got: ", words)
		return "run " + val
	}
	return "No such command!"
}
