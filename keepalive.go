package main

import (
	"fmt"
	"strings"

	"github.com/ateliersjp/http"
)

func disableKeepAlive(m *http.Msg) {
	for i, l := range m.Headers {
		if i == 0 {
			continue
		}

		if k, v, ok := strings.Cut(l, ":"); ok {
			k = k[strings.LastIndex(k, "-") + 1:]
			v = strings.TrimSpace(v)

			if strings.EqualFold(k, "connection") {
				v = strings.Replace(v, "keep-alive", "close", 1)
				v = strings.Replace(v, "Keep-Alive", "Close", 1)
				m.Headers[i] = fmt.Sprintf("%s: %s", k, v)
				break
			}
		}
	}
}
