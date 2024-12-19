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
	Hostname   string
	TLS        bool
}

func getDefaultPort(host *Host) string {
	if host.TLS || host.Connect {
		return ":443"
	}

	return":80"
}

func NewRequest(conn net.Conn) *Request {
	m, err := http.ReadMsg(conn)

	if err != nil || len(m.Headers) == 0 {
		return nil
	}

	host, ok := getHost(m)
	if !ok {
		return nil
	}

	r := &Request{}
	r.conn = conn

	if strings.HasSuffix(m.Headers[0], " HTTP/1.1") {
		disableKeepAlive(m)
	}

	r.TLS = host.TLS
	if tls, ok := getSchemeByHeaders(m); ok {
		r.TLS = tls
	}

	setHostname(m, host.Hostname)

	if strings.Contains(host.Hostname, ":") {
		r.Hostname = host.Hostname
	} else {
		port := getDefaultPort(host)
		r.Hostname = fmt.Sprintf("%s%s", host.Hostname, port)
	}

	if host.Connect {
		r.Reader = m.Body
		return r.OK()
	}

	r.Reader = m.Reader()
	return r
}

func (r *Request) OK() *Request {
	_, err := fmt.Fprintf(r.conn, "HTTP/1.1 200 OK\r\n\r\n")

	if err != nil {
		return nil
	}

	return r
}

func (r *Request) Dial() (net.Conn, error) {
	if r.TLS {
		return tls.Dial(PROTOCOL, r.Hostname, &tls.Config{})
	} else {
		return net.Dial(PROTOCOL, r.Hostname)
	}
}
