package webrtc

import (
  "io"

  "github.com/pion/webrtc/v4"
)

type peer struct {
  name  string
  class byte           // otm, mtm lct...
  stmt  byte           // stream type: "rtpc/weprtc"
  sigw  chan io.Writer // write ice to peer
  stmr  stmReader
  stmw  stmWriter
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
func newPeer(sigclass, stmTyp byte, name string) (*peer, error) {
  p := &peer{stmt: stmTyp, class: sigclass, name: name}
  switch sigclass {
  case SigClassOTOCaller:
  case SigClassOTOCallee:
  case SigClassAllocOTM:
  case SigClassAllocMTM:
  case SigClassAllocLCT:
  }

  return p, nil
}

type webrtcReader struct {
  tracks []*webrtc.TrackRemote
}

type webWriter struct {
  tracks []*webrtc.TrackLocal
}
