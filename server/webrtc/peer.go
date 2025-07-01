package webrtc

import (
  "io"

  "github.com/pion/webrtc/v4"
)

type peer struct {
  name string
  sigw chan io.Writer // write ice to peer
  stmt byte           // stream type: "rtpc/weprtc"
  stmr stmReader
  stmw stmWriter
}

type stmData struct {
  stmType byte
  video   []byte
  audio   []byte
}

type stmReader interface {
  Read() (*stmData, error)
}

type stmWriter interface {
  Write(d *stmData) error
}

// newPeer ...
func newPeer(stmTyp byte) (*peer, error) {
  p := &peer{stmt: stmTyp}
  return p, nil
}

type webrtcReader struct {
  tracks []*webrtc.TrackRemote
}

type webWriter struct {
  tracks []*webrtc.TrackLocal
}
