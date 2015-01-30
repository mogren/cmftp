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
	log.Print("hej", *verbose)
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
		io.WriteString(c, "Invalid Username\n")
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
	client.WriteLinesFrom(client.channel)
}

func promtLogin(c net.Conn, bufc *bufio.Reader) string {
	io.WriteString(c, "220 FTP Server ready.\n")
	io.WriteString(c, "Username: ")
	username, err := bufc.ReadString('\n')
	if err == nil {
		io.WriteString(c, "331 User "+username+" OK. Password required")
		passwd, _ := bufc.ReadString('\n')
		log.Println("Passwd: " + passwd)
	}
	return string(username)
}

func logMessages(logChan <-chan string) {
	for msg := range logChan {
		log.Printf("client msg : %s", msg)
	}
}
