package webrtc

type otochan struct {
  peers [2]*peer
}

type otmroom struct {
  name     string
  self     *peer
  watchers []*peer
}

type mtmroom struct {
  name string
  pct  []*peer
}

type lctroom struct {
  name string
  tch  *peer
  stu  []*peer
}

type group interface {
  hSig(m *msg) error
  hData() // track and data channel
}
