package webrtc

import (
  "github.com/yolksys/emei/env"
  "github.com/gorilla/websocket"
)

type Webrtc struct{}

// called signal server by mdn
func (w *Webrtc) Signal(e env.Env, conn *websocket.Conn) {
}

type msg struct {
  Class  byte   `json:"class,omitempty"` // alloc/webrtc/rtp
  Typ    byte   `json:"type,omitempty"`  // oto.webrtc/otm.
  Name   string `json:"name,omitempty"`
  Target string `json:"target,omitempty"`
  Cont   string `json:"cont,omitempty"` // content
}

type client struct {
  name   string
  oto    [][2]*peer
  otm    *peer
  mtm    *peer
  watchs []*peer
}
