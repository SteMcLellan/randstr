package randstr

import(
    "math/rand"
)

type StringSource interface {
  Rune() rune
  Random(src *rand.Rand)
}

type stringSource struct {
  random *rand.Rand
}

func (ss *stringSource) Rune() rune {
  chr := ss.random.Intn(52) + 65
  if chr > 90 {
    chr = chr + 6
  }

  return rune(chr)
}

func (ss *stringSource) Random(rnd *rand.Rand) {
  ss.random = rnd
}

type RandString struct {
  strSrc StringSource
}

func New(rng *rand.Rand) *RandString {
  return &RandString{rng}
}

func (rnd *RandString) Next() rune {
  var n rune
  n = rnd.strSrc.Rune()

  return n
}

func (rnd *RandString) NextStr(l int) string {
  var n []rune
  n = make([]rune, l)

  for i := range n {
    n[i] = rnd.Next()
  }

  return string(n)
}

func (rnd *RandString) Read(b []byte) (n int, err error) {
  for i := range b {
    b[i] = byte(rnd.Next())
  }

  return len(b), nil
}
