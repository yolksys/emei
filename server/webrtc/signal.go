package webrtc

import (
  "encoding/json"
  "fmt"
  "io"
  "reflect"

  "github.com/gorilla/websocket"
  "github.com/yolksys/emei/env"
)

type Webrtc0 struct{}

// called signal server by mdn
func (w *Webrtc0) Signal(e env.Env, conn any) {
  defer e.Return()

  do_ := func(m *msg) {
    switch m.Class {
    case SigClassOTOCaller:
    case SigClassOTOCallee:
    case SigClassAllocOTM:
    case SigClassAllocMTM:
    case SigClassAllocLCT:
    }
  }

  var m msg
  switch c := conn.(type) {
  case io.ReadWriter:
    dec := json.NewDecoder(c)
    for {
      err := dec.Decode(&m)
      e.AssertErr(err)
      do_(&m)
    }

  case *websocket.Conn:
    for {
      err := c.ReadJSON(&m)
      e.AssertErr(err)
      do_(&m)
    }
  default:
    e.AssertErr(fmt.Errorf("fail:signal, reason: err conn type, conn type:%+v", reflect.TypeOf(conn).Name()))
  }
}

type msg struct {
  Class  byte   `json:"class,omitempty"` // alloc/webrtc/rtp
  Typ    byte   `json:"type,omitempty"`  // oto.webrtc/otm.
  Name   string `json:"name,omitempty"`
  Target string `json:"target,omitempty"`

  Attrs map[string]string `json:"attrs,omitempty"` // content
}

type client struct {
  name string
  oto  []*otochan // two peer all are webrtc then pass data by turn
  otm  map[string]*otmroom
  mtm  map[string]*mtmroom
}

var (
  _peers   = map[string]*peer{} // key==peername
  _rooms   = map[string]group{} // key==peername, value= *otochan, *otmroom, *mtmroom, *lctroom
  _clients = map[string]*client{}
  // _oto   = map[string]*otochan{}
  // _otm   = map[string]*otmroom{}
  // _mtm   = map[string]*mtmroom{}
  // _lct   = map[string]*lctroom{}
)
