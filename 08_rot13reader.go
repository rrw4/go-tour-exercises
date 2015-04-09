package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		if (b[i] >= 65 && b[i] <= 77) || (b[i] >= 97 && b[i] <= 109) {
			b[i] = b[i] + 13
		} else if (b[i] >= 78 && b[i] <= 90) || (b[i] >= 110 && b[i] <= 122) {
			b[i] = b[i] - 13
		}
	}

	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
