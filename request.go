package main

import (
	"fmt"
	"io"
	"net"
	"crypto/tls"
	"strings"

	"github.com/ateliersjp/http"
)

type Request struct {
	io.Reader
	conn    net.Conn
	host    string
	proto   string
	tls     bool
}

func NewRequest(conn net.Conn) *Request {
	m, err := http.ReadMsg(conn)
	if err != nil {
		return nil
	}
	host, proto, tls, tun := getHostname(m)
	if host == "" {
		return nil
	}
	if proto == "HTTP/1.1" {
		disableKeepAlive(m)
	} else if proto != "HTTP/1.0" {
		return nil
	}
	setHostname(m, host)
	if !strings.Contains(host, ":") {
		if tls || tun {
			host = fmt.Sprintf("%s:443", host)
		} else {
			host = fmt.Sprintf("%s:80", host)
		}
	}
	r := &Request{}
	r.conn, r.host, r.proto, r.tls = conn, host, proto, tls
	if tun {
		r.Reader = m.Body
		return r.OK()
	}
	r.Reader = m.Reader()
	return r
}

func (r *Request) OK() *Request {
	_, err := fmt.Fprintf(r.conn, "%s 200 OK\r\n\r\n", r.proto)
	if err != nil {
		return nil
	}
	return r
}

func (r *Request) Dial() (net.Conn, error) {
	if r.tls {
		return tls.Dial(PROTOCOL, r.host, &tls.Config{})
	} else {
		return net.Dial(PROTOCOL, r.host)
	}
}
