package main

import (
	"flag"
)

var (
	PROTOCOL = "tcp"
	PORT uint
	ROOT_DOMAIN string
	SCHEME_HEADER string
	USER_AGENT string
)

func init() {
	flag.UintVar(&PORT, "p", 1080, "listening port")
	flag.StringVar(&ROOT_DOMAIN, "d", "", "root domain for subdomain-based URIs")
	flag.StringVar(&SCHEME_HEADER, "H", "X-Forwarded-Proto", "request header for identifying the scheme")
	flag.StringVar(&USER_AGENT, "A", "", "replace the client's user agent with it")

	flag.Parse()
}
