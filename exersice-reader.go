package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add the "func (T) Read([]byte) (int, err) {}" into MyReader struct
func (r MyReader) Read(b []byte) (int, error) {
	for n := range b {
		b[n] = 'A'	
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}