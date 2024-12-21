package main

import (
	"flag"
)

var (
	PROTOCOL = "tcp"
	PORT uint
)

func init() {
	flag.UintVar(&PORT, "p", 1080, "local port number")
}
