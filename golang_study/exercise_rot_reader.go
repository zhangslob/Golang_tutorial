package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (n int, e error) {
	n, e = rr.r.Read(b)
	if e != nil {
		return 0, e
	}
	for i := range b {
		if b[i] == ' ' {
			continue
		}
		switch {
		case 'A' <= b[i] && b[i] < 'N':
			b[i] += 13
		case 'a' <= b[i] && b[i] < 'n':
			b[i] += 13
		default:
			b[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, _ = io.Copy(os.Stdout, &r)
}
