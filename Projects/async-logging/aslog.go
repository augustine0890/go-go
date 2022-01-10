// Package aslog provides a simple asynchronous logger that will write to provided io.Writers without blocking calling
package aslog

import (
	"io"
	"os"
	"sync"
)

// Aslog is a type that defines a logger. It can be used to write log messages synchronously (via the Write method)
// or asynchrouusly via the channel returned by the MessageChannel accessor.
type Alog struct {
	dest               io.Writer
	m                  *sync.Mutex
	msgCh              chan string
	errorCh            chan error
	shutdownCh         chan struct{}
	shutdownCompleteCh chan struct{}
}

// New creates a new Alog object that writes to the provided io.Writer
// If nil is provided the output will be directed to os.Stdout
func New(w io.Writer) *Alog {
	if w == nil {
		w = os.Stdout
	}
	return &Alog{
		dest: w,
	}
}
