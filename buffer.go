// +build !darwin

package clip

import (
	"fmt"
	"os"
)

type Clipboard struct {
	buf []byte
}

func init() {
	fmt.Fprintln(os.Stderr, "clipboard: falling back to internal buffer")
}

func (c *Clipboard) Get() ([]byte, error) {
	buf := make([]byte, len(c.buf))
	copy(buf, c.buf)
	return buf, nil
}

func (c *Clipboard) Put(buf []byte) error {
	if cap(c.buf) < len(buf) {
		c.buf = make([]byte, len(buf))
	}
	c.buf = c.buf[:len(buf)]
	copy(c.buf, buf)
	return nil
}
