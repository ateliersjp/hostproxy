package main

import (
	"fmt"
	"strings"

	"github.com/ateliersjp/http"
)

func getHostname(m *http.Msg) (host, port string) {
	for _, l := range m.Headers {
		if k, v, ok := strings.Cut(l, ":"); ok {
			if strings.EqualFold(k, "Host") {
				host = strings.TrimSpace(v)
				if n := strings.Index(host, ":"); n != -1 {
					host, port = host[:n], host[n:]
				}

				break
			}
		}
	}
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
