package main

import (
	"fmt"
	"strings"

	"go4.org/strutil"
	"github.com/ateliersjp/http"
)

func parseHostname(args []string) (host, proto string, tls, tun bool) {
	if len(args) != 3 {
		return
	}
	proto = args[2]
	if args[0] == "CONNECT" {
		host, tun = args[1], true
		return
	}
	uri, _ := strings.CutPrefix(args[1], "/")
	if uri, tls = strings.CutPrefix(uri, "https://"); !tls {
		uri, _ = strings.CutPrefix(uri, "http://")
	}
	host, uri, _ = strings.Cut(uri, "/")
	args[1] = fmt.Sprintf("/%v", uri)
	return
}

func getHostname(m *http.Msg) (host, proto string, tls, tun bool) {
	if len(m.Headers) > 0 {
		args := strings.SplitN(m.Headers[0], " ", 3)
		host, proto, tls, tun = parseHostname(args)
		m.Headers[0] = strings.Join(args, " ")
	}
	return
}

func setHostname(m *http.Msg, host string) {
	hh := fmt.Sprintf("Host: %s", host)
	for i, line := range m.Headers {
		if strutil.HasPrefixFold(line, "Host:") {
			m.Headers[i] = hh
			return
		}
	}
	m.Headers = append(m.Headers, hh)
}
