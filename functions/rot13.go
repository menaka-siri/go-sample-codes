package functions

import (
	"io"
)

type Rot13Reader struct {
	R io.Reader
}

func (rot *Rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.R.Read(b)
	if err == io.EOF {
		return 0, err
	}

	for x, v := range b {
		b[x] = byte(rot13(v))
	}
	return n, err
}

func rot13(x byte) byte {
	capital := x >= 'A' && x <= 'Z'
	if !capital && (x < 'a' || x > 'z') {
		return x // Not a letter
	}

	x += 13
	if capital && x > 'Z' || !capital && x > 'z' {
		x -= 26
	}
	return x
}
