package randstr

import (
  "testing"
)

func TestRandomWriter(t *testing.T) {
  b := make([]byte, 10)
  var n, _ = Read(b)

  if(n != 10) {
    t.Errorf("Buffer not the right length")
  }
}
