package main

import (
	"io"
	"sync"
)

type WaitGroup struct {
	sync.WaitGroup
}

func NewWaitGroup() *WaitGroup {
	var wg WaitGroup
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

func (wg *WaitGroup) Copy(dst io.Writer, src io.Reader) (int64, error) {
	defer wg.Done()
	defer closeWrite(dst)
	return io.Copy(dst, src)
}
