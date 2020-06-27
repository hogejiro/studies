package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13r rot13Reader) Read(buff []byte) (n int, err error) {
	n, err = r13r.r.Read(buff)
	if err == nil {
		for i := range buff {
			buff[i] = DecodeRot13(buff[i])
		}
	}
	return
}

func DecodeRot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return (b-'A'+13)%26 + 'A'
	case 'a' <= b && b <= 'z':
		return (b-'a'+13)%26 + 'a'
	default:
		return b
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
