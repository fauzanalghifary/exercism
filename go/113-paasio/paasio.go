package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	sync.RWMutex
	r    io.Reader
	n    int64
	nops int
}
type writeCounter struct {
	sync.RWMutex
	w    io.Writer
	n    int64
	nops int
}

// For the return of the function NewReadWriteCounter,
// you must also define a type that satisfies the ReadWriteCounter interface.
type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{w: writer}
}
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{r: reader}
}
func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewReadCounter(readwriter), NewWriteCounter(readwriter)}
}
func (rc *readCounter) Read(p []byte) (int, error) {
	rc.Lock()
	defer rc.Unlock()
	n, err := rc.r.Read(p)
	rc.nops++
	rc.n += int64(n)
	return n, err
}
func (rc *readCounter) ReadCount() (int64, int) {
	rc.Lock()
	defer rc.Unlock()
	return rc.n, rc.nops
}
func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.Lock()
	defer wc.Unlock()
	n, err := wc.w.Write(p)
	wc.nops++
	wc.n += int64(n)
	return n, err
}
func (wc *writeCounter) WriteCount() (int64, int) {
	wc.Lock()
	defer wc.Unlock()
	return wc.n, wc.nops
}
