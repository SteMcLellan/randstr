package randstr

import (
	"testing"
)

func TestNext(t *testing.T) {
	r := New(NewSource(1))
	c := r.Next()

	if c != 'x' {
		t.Errorf("Incorect letter (%v) returned, expected x.", c, string(c))
	}
}

func TestReader(t *testing.T) {
	r := New(NewSource(1))
	var buf = make([]byte, 10)

	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("Error reading from random string stream.")
	}

	if n != 10 {
		t.Errorf("Did not read the full length of buffer (10)")
	}

	str := string(buf)
	if str != "xvLbZGBAIc" {
		t.Errorf("Buffer reads as %s", string(buf))
	}
}
