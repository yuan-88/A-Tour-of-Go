package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Method of io.Reader:  func (T) Read([]byte) (int, error) {}
func (rot rot13Reader) Read(b []byte) (int, error) {
	// Read strings of rot
	n, err := rot.r.Read(b)
	// Convert char of string from rot that are recorded in b []byte
	for i, char := range b {
		switch {
		case char >= 'A' && char < 'N', char <= 'a' && char > 'n':
			b[i] += 13
		case char >= 'N' && char <= 'Z', char >= 'n' && char <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}