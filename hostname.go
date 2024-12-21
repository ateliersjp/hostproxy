package main

import (
	"fmt"
	"strings"

	"github.com/ateliersjp/http"
)

func getHostname(m *http.Msg) (host, proto string, tls, tun bool) {
	if len(m.Headers) == 0 {
		return
	}

	args := strings.SplitN(m.Headers[0], " ", 3)

	if len(args) != 3 {
		return
	}

	proto = args[2]

	if args[0] == "CONNECT" {
		host, tun = args[1], true
		return
	}

	uri := strings.TrimPrefix(args[1], "/")

	uri, tls = strings.CutPrefix(uri, "https://")

	if !tls {
		uri, _ = strings.CutPrefix(uri, "http://")
	}

	host, uri, _ = strings.Cut(uri, "/")
	args[1] = fmt.Sprintf("/%v", uri)

	m.Headers[0] = strings.Join(args, " ")

	return
}

func setHostname(m *http.Msg, host string) {
	for i, l := range m.Headers {
		if k, _, ok := strings.Cut(l, ":"); ok {
			if strings.EqualFold(k, "Host") {
				m.Headers[i] = fmt.Sprintf("%s: %s", k, host)
				break
			}
		}
	}
	return
}
