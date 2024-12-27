package main

import (
	"io"
	"sync"
)

type waitGroup struct {
	sync.WaitGroup
}

func newWaitGroup() *waitGroup {
	var wg waitGroup
	wg.Add(2)
	return &wg
}

type closeWriter interface {
	CloseWrite() error
} 

func closeWrite(w io.Writer) {
	if conn, ok := w.(closeWriter); ok {
		conn.CloseWrite()
	}
}

func (wg *waitGroup) copy(dst io.Writer, src io.Reader) (int64, error) {
	defer wg.Done()
	defer closeWrite(dst)
	return io.Copy(dst, src)
}
