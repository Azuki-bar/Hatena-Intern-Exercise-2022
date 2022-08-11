package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	buf := make([]byte, len(b))
	n, err = rot.r.Read(buf)
	if err != nil {
		return n, err
	}
	var i int
	for i = 0; i < n; i++ {
		elem := buf[i]
		switch {
		case 'A' <= elem && elem <= 'Z':
			b[i] = (elem-'A'+13)%26 + 'A'
		case 'a' <= elem && elem <= 'z':
			b[i] = (elem-'a'+13)%26 + 'a'
		case elem == 0:
			return n, io.EOF
		default:
			b[i] = elem
		}
	}
	return i, nil
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
