package main

import (
	"fmt"
	"strings"

	"github.com/ateliersjp/http"
	texttransform "github.com/tenntenn/text/transform"
)

type config struct {
	host    string
	proto   string
	tls     bool
	tun     bool
	repl    texttransform.ReplaceByteTable
}

func newConfig(m *http.Msg) (conf *config, ok bool) {
	if len(m.Headers) == 0 {
		return
	}

	args := strings.SplitN(m.Headers[0], " ", 3)

	if len(args) != 3 {
		return
	}

	conf, ok = &config{}, true
	conf.proto = args[2]

	hostname, found := removeHeader(m, "Host")

	if args[0] == "CONNECT" {
		conf.host, conf.tun = args[1], true
		return
	}

	if found && len(ROOT_DOMAIN) != 0 {
		var port string
		if n := strings.LastIndex(hostname, ":"); n != -1 {
			hostname, port = hostname[:n], hostname[n:]
		}

		suffix := fmt.Sprintf(".%s", ROOT_DOMAIN)
		if hostname, found = strings.TrimSuffix(hostname, suffix); found {
			conf.repl = append(conf.repl, []byte(hostname))
			hostname = hostname[:len(hostname)-len(suffix)]
			hostname = strings.ReplaceAll(hostname, "--", ".")
			conf.repl = append(conf.repl, []byte(hostname))
			conf.host = fmt.Sprintf("%s%s", hostname, port)

			if scheme, _ := removeHeader(m, SCHEME_HEADER); scheme == "https" || scheme == "on" {
				conf.tls = true
			}

			return
		}
	}

	uri := strings.TrimPrefix(args[1], "/")

	if uri, found = strings.CutPrefix(uri, "https://"); found {
		conf.tls = true
	} else if uri, found = strings.CutPrefix(uri, "http://"); found {
		conf.tls = false
	} else if scheme, _ := removeHeader(m, SCHEME_HEADER); scheme == "https" || scheme == "on" {
		conf.tls = true
	}

	conf.host, uri, _ = strings.Cut(uri, "/")
	args[1] = fmt.Sprintf("/%v", uri)

	m.Headers[0] = strings.Join(args, " ")

	return
}
