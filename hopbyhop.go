package main

import (
	"strings"

	"github.com/ateliersjp/http"
)

func removeHopByHopHeaders(m *http.Msg, names string) {
	for _, name := range strings.Split(names, ",") {
		removeHeader(m, strings.TrimSpace(name))
	}
}
