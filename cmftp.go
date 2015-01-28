package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		port    = flag.Int("port", 20000, "FTP server port. (1024 – 49151)")
		verbose = flag.Bool("v", false, "Write log messages to stdout")
	)
	flag.Parse()

	if *port < 1024 || *port > 49151 {
		flag.Usage()
		os.Exit(1)
	}
	if *verbose {
		fmt.Println("port: ", *port)
	}
	log.Print("hej", *verbose)
}
