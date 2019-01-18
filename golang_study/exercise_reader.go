package main

import "golang.org/x/tour/reader"

type MyReader struct {}

func (r MyReader) Read(s []byte) (n int, err error) {
	s = s[:cap(s)]
	for i := range s{
		s[i] = 'A'
	}
	return cap(s), nil
}

func main() {
	reader.Validate(MyReader{})
}
