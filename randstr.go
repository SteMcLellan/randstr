package randstr

import(
    "math/rand"
)

type StringSource interface {
  Rune() rune
}

type stringSource struct {
  random *rand.Rand
}

func (ss stringSource) Rune() rune {
  chr := ss.random.Intn(52) + 65
  if chr > 90 {
    chr = chr + 6
  }

  return rune(chr)
}

type RandString struct {
  strSrc StringSource
}

func NewSource(seed int64) StringSource {
  randSource := rand.NewSource(seed)
  rand := rand.New(randSource)
  ss := &stringSource{rand}

  return ss
}

func New(ss StringSource) *RandString {
  return &RandString{ss}
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
