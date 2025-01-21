package main

import (
	"strings"

	"github.com/ateliersjp/http"
)

func removeHeader(m *http.Msg, n string) (v string, ok bool) {
	for i, l := range m.Headers {
		if i == 0 {
			continue
		}

		var k string
		if k, v, ok = strings.Cut(l, ":"); ok {
			if ok = strings.EqualFold(k, n); ok {
				m.Headers = m.Headers[:i+copy(m.Headers[i:], m.Headers[i+1:])]
				v = strings.TrimSpace(v)
				break
			}
		}
	}

	return
}
