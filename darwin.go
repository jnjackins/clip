// +build darwin

package clip

import (
	"bytes"
	"os/exec"
)

type Clipboard struct {
}

func (c *Clipboard) Get() ([]byte, error) {
	cmd := exec.Command("pbpaste")
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	err := cmd.Run()
	return buf.Bytes(), err
}

func (c *Clipboard) Put(buf []byte) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = bytes.NewBuffer(buf)
	return cmd.Run()
}
