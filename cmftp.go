package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const ADMIN_NAME string = "admin"
const ADMIN_PWD string = "password"

func main() {
	var (
		iport   = flag.Int("port", 2121, "FTP server port. (1024 – 49151)")
		verbose = flag.Bool("v", false, "Write log messages to stdout")
	)
	flag.Parse()
	if *iport < 1024 || *iport > 49151 {
		flag.Usage()
		os.Exit(1)
	}
	port := ":" + strconv.Itoa(*iport)
	if *verbose {
		fmt.Println("port: ", port)
	}
	log.Println("cmftp started. (Verbose: ", *verbose, ")")
	// Logger
	logChan := make(chan string)
	go logMessages(logChan)
	// Listen
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn, logChan)
	}
}

func handleConnection(c net.Conn, logChan chan<- string) {
	bufc := bufio.NewReader(c)
	defer c.Close()

	client := Client{
		connection: c,
		username:   promtLogin(c, bufc),
		channel:    make(chan string),
		admin:      true,
	}
	if strings.TrimSpace(client.username) == "" {
		return
	}

	// Register user
	//addchan <- client
	defer func() {
		logChan <- fmt.Sprintf("User %s left the chat room.\n", client.username)
		log.Printf("Connection from %v closed.\n", c.RemoteAddr())
		//rmchan <- client
	}()
	io.WriteString(c, fmt.Sprintf("230 Welcome to cmftp, %s!\n\n", client.username))
	logChan <- fmt.Sprintf("User %s has authenticated.\n", client.username)

	// I/O
	go client.ReadLinesInto(logChan)
	// Echo back
	client.WriteLinesFrom(client.channel)
}

func promtLogin(c net.Conn, bufc *bufio.Reader) string {
	io.WriteString(c, "220 FTP Server ready.\n")
	io.WriteString(c, "Username: ")
	username, _, err := bufc.ReadLine()
	ret := string(username)
	if err == nil {
		// TODO: check username in user list
		if ret != ADMIN_NAME {
			io.WriteString(c, "332 Username "+ret+" not found\n")
			return ""
		} else {
			io.WriteString(c, "331 User "+ret+" OK. Password required\n")
			passwd, _, err := bufc.ReadLine()
			log.Println("Passwd: " + string(passwd))
			if err != nil || string(passwd) != ADMIN_PWD {
				io.WriteString(c, "530 Not logged in.\n")
				return ""
			}
		}
	}
	log.Println("ret: " + ret)
	return ret
}

func logMessages(logChan <-chan string) {
	for msg := range logChan {
		log.Printf("%s", msg)
	}
}
