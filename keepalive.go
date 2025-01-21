package main

import (
	"fmt"
	"strings"

	"github.com/ateliersjp/http"
)

func disableKeepAlive(m *http.Msg, name string) {
	if v, ok := removeHeader(m, name); ok {
		if !strings.EqualFold(v, "close") {
			removeHopByHopHeaders(m, v)
		}

		m.Headers = append(m.Headers, fmt.Sprintf("%s: close", name))
	}
}
