package main

import (
	"strings"

	"github.com/ateliersjp/http"
)

func getSchemeByHeaders(m *http.Msg) (tls, ok bool) {
	for _, l := range m.Headers {
		if k, v, ok := strings.Cut(l, ":"); ok {
			n := strings.LastIndex(k, "-") + 1
			k = k[n:]
			v = strings.TrimSpace(v)

			if strings.EqualFold(k, "proto") || strings.EqualFold(k, "protocol") || strings.EqualFold(k, "scheme") {
				if v == "https" {
					tls = true
					ok = true
					break
				}

				if v == "http" {
					tls = false
					ok = true
					break
				}
			}

			if strings.EqualFold(k, "https") || strings.EqualFold(k, "tls") {
				if v == "on" {
					tls = true
					ok = true
					break
				}

				if v == "off" {
					tls = false
					ok = true
					break
				}
			}
		}
	}

	return
}

func getSchemeByURI(uri *string) (tls, ok bool) {
	*uri, tls = strings.CutPrefix(*uri, "https://")
	*uri, ok = strings.CutPrefix(*uri, "http://")
	return tls, tls || ok
}
