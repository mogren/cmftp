package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

// Client is a connected user to the server
type Client struct {
	connection net.Conn
	username   string
	homeDir    string
	channel    chan string
	admin      bool
}

// ReadLinesInto will read lines from the Client socket
func (c Client) ReadLinesInto(ch chan<- string) {
	bufc := bufio.NewReader(c.connection)
	for {
		line, _, err := bufc.ReadLine()
		if err != nil {
			break
		}
		//fmt.Printf("debug input: %v, %s, %x \n", line, line, line)
		str = string(line)
		if str != "" && (int(line[0]) == 4 || strings.HasPrefix(str, "/quit")) {
			c.connection.Close()
			return
		}
		cmd := strings.TrimSpace(str)
		if cmd != "" {
			// Execute command?
			result := RunCommand(cmd)
			fmt.Println(result)
			ch <- fmt.Sprintf("%s: %s", c.username, str)
		}
	}
}

// WriteLinesFrom will echo back to the client
func (c Client) WriteLinesFrom(ch <-chan string) {
	for msg := range ch {
		_, err := io.WriteString(c.connection, msg)
		if err != nil {
			return
		}
	}
}
