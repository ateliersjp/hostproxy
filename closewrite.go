package main

import (
	"io"
)

type closeWriter interface {
	CloseWrite() error
} 

func closeWrite(w io.Writer) {
	if conn, ok := w.(closeWriter); ok {
		conn.CloseWrite()
	}
}
