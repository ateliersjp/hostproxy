package main

import (
	"fmt"
	"io"
	"net"
	"crypto/tls"
	"strings"

	"github.com/ateliersjp/http"
	texttransform "github.com/tenntenn/text/transform"
)

type request struct {
	io.Reader
	config
	conn    net.Conn
}

func newRequest(conn net.Conn) *request {
	m, err := http.ReadMsg(conn)
	if err != nil {
		return nil
	}

	conf, ok := newConfig(m)
	if !ok {
		return nil
	}

	switch conf.proto {
	case "HTTP/1.1":
		disableKeepAlive(m, "Connection")
		disableKeepAlive(m, "Proxy-Connection")
	case "HTTP/1.0":
	default:
		return nil
	}

	removeHopByHopHeaders(m, "Keep-Alive, Transfer-Encoding, TE, Connection, Trailer, Upgrade, Proxy-Authorization, Proxy-Authenticate, Forwarded, X-Forwarded-Host, X-Forwarded-For, X-Real-IP")

	if len(USER_AGENT) != 0 {
		removeHeader(m, "User-Agent")
		ua := fmt.Sprintf("User-Agent: %s", USER_AGENT)
		m.Headers = append(m.Headers, ua)
	}

	m.Headers = append(m.Headers, fmt.Sprintf("Host: %s", conf.host))

	if !strings.Contains(conf.host, ":") {
		if conf.tls || conf.tun {
			conf.host = fmt.Sprintf("%s:443", conf.host)
		} else {
			conf.host = fmt.Sprintf("%s:80", conf.host)
		}
	}

	r := &request{}
	r.conn, r.config = conn, *conf

	if conf.tun {
		r.Reader = m.Body
		return r.ok()
	}

	if len(conf.repl) != 0 {
		t := texttransform.ReplaceAll(conf.repl)
		if mod, err := m.Transform(t); err == nil {
			m = mod
		}
	}

	r.Reader = m.Reader()
	return r
}

func (r *request) ok() *request {
	_, err := fmt.Fprintf(r.conn, "%s 200 OK\r\n\r\n", r.proto)

	if err != nil {
		return nil
	}

	return r
}

func (r *request) dial() (net.Conn, error) {
	if r.tls {
		return tls.Dial(PROTOCOL, r.host, &tls.Config{})
	} else {
		return net.Dial(PROTOCOL, r.host)
	}
}
