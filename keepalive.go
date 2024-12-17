package main

import (
	"strings"

	"go4.org/strutil"
	"github.com/ateliersjp/http"
)

func disableKeepAlive(m *http.Msg) {
	for i, line := range m.Headers {
		if strutil.HasPrefixFold(line, "Connection:") || strutil.ContainsFold(line, "-Connection:") {
			line = strings.Replace(line, "keep-alive", "close", 1)
			line = strings.Replace(line, "Keep-Alive", "Close", 1)
			m.Headers[i] = line
			break
		}
	}
}
