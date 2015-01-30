package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

type Client struct {
	connection net.Conn
	username   string
	homeDir    string
	channel    chan string
	admin      bool
}

func (c Client) ReadLinesInto(ch chan<- string) {
	bufc := bufio.NewReader(c.connection)
	for {
		line, err := bufc.ReadString('\n')
		if err != nil {
			break
		}
		if strings.HasPrefix(line, "/quit") {
			c.connection.Close()
			return
		}
		ch <- fmt.Sprintf("%s: %s", c.username, line)
	}
}

func (c Client) WriteLinesFrom(ch <-chan string) {
	for msg := range ch {
		_, err := io.WriteString(c.connection, msg)
		if err != nil {
			return
		}
	}
}
