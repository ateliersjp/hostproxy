package main

import (
	"fmt"
	"strings"

	"github.com/ateliersjp/http"
)

type Host struct {
	Hostname   string
	TLS        bool
	Connect    bool
}

func parseHost(m *http.Msg) (host *Host, path string) {
	if method, uri, ok := strings.Cut(m.Headers[0], " "); ok && method != "CONNECT" {
		if uri, _, ok = strings.Cut(uri, " "); ok {
			host = &Host{}
			uri = strings.TrimPrefix(uri, "/")
		
			if tls, ok := getSchemeByHeaders(m); ok {
				host.TLS = tls
			} else if tls, ok := getSchemeByURI(&uri); ok {
				host.TLS = tls
			}
		
			host.Hostname, path, _ = strings.Cut(uri, "/")
		}
	}

	return
}

func parseConnectHost(m *http.Msg) (host *Host, ok bool) {
	if method, uri, found := strings.Cut(m.Headers[0], " "); found && method == "CONNECT" {
		if uri, _, ok = strings.Cut(uri, " "); ok {
			host = &Host{}
			host.Hostname = uri
			host.Connect = true
		}
	}

	return
}

func getHostByHeaders(m *http.Msg) (host *Host, ok bool) {
	if DOMAIN == "" {
		return
	}

	hostname, port := getHostname(m)

	if hostname, ok = strings.CutSuffix(hostname, fmt.Sprintf(".%s", DOMAIN)); ok {
		host = &Host{}
		host.Hostname = fmt.Sprintf("%s%s", hostname, port)
	}

	return
}

func getHostByRequestLine(m *http.Msg) (host *Host, ok bool) {
	args := strings.SplitN(m.Headers[0], " ", 3)

	if len(args) != 3 {
		return
	}

	if host, ok := parseConnectHost(m); ok {
		return host, true
	}

	host, path := parseHost(m)
	args[1] = fmt.Sprintf("/%v", path)
	m.Headers[0] = strings.Join(args, " ")
	return host, true
}

func getHost(m *http.Msg) (host *Host, ok bool) {
	host, ok = getHostByHeaders(m)

	if ok {
		return
	}

	host, ok = getHostByRequestLine(m)
	return
}
