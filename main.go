package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"io"
	"os"
)

func main() {
	flag.Parse()
	ln, err := net.Listen(PROTOCOL, fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	if req := NewRequest(conn); req != nil {
		if remote, err := req.Dial(); err == nil {
			defer remote.Close()
			wg := NewWaitGroup()
			go wg.Copy(remote, io.TeeReader(req, os.Stdout))
			go wg.Copy(conn, remote)
			wg.Wait()
		}
	}
}
