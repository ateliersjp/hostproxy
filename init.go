package main

import (
	"flag"
)

var (
	PROTOCOL = "tcp"
	PORT uint
	DOMAIN string
)

func init() {
	flag.UintVar(&PORT, "p", 1080, "local port number")
	flag.StringVar(&DOMAIN, "d", "", "domain of the proxy")
}
