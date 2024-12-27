package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
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
	if req := newRequest(conn); req != nil {
		if remote, err := req.dial(); err == nil {
			defer remote.Close()
			wg := newWaitGroup()
			go wg.copy(remote, req)
			go wg.copy(conn, remote)
			wg.Wait()
		}
	}
}
