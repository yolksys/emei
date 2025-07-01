package web

import (
  "errors"
  "fmt"
  "mime/multipart"
  "net/http"
  "reflect"
  "strings"

  "github.com/yolksys/emei/env"
  "github.com/yolksys/emei/log"
  "github.com/yolksys/emei/rpc/coder"
  "github.com/yolksys/emei/utils"
)

// route ...
func (s *webrpc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // buf := utils.NewBuffer()
  // io.Copy(buf, r.Body)b
  // r.Body.Close()
  cntTyp := r.Header.Get("Content-Type")
  io_ := r.Body
  stream := [1]any{}
  var err error
  if cntTyp == "multipart/form-data" {
    defer r.Body.Close()
    err = r.ParseMultipartForm(1024 * 1024 * 10)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      sss := fmt.Sprintf("fail:http.ParseMultipartForm, reason:%s, path:%s", err.Error(), r.URL.Path)
      w.Write([]byte(sss))
      log.Error("msg", sss)
      return
    }
  }

  p := strings.Trim(r.URL.Path, "/")
  e := env.New("web", strings.ReplaceAll(p, "/", "."), newEnc(w), newDec(io_))
  defer e.Finish()
  e.Assert()
  wh_ := func() { w.WriteHeader(http.StatusInternalServerError) }

  if r.Header.Get("upgrade") == "websocket" {
    c, err := newWebSock(w, r)
    e.AssertErr(err, wh_)
    stream[0] = c
  } else if cntTyp == "multipart/form-data" {
    defer func() {
      for _, v := range stream {
        v.(multipart.File).Close()
      }
    }()

    io_ = utils.NewRbuffer([]byte(r.Form["params"][0]))
    if len(r.MultipartForm.File["file"]) == 0 {
      e.AssertErr(fmt.Errorf("fail:file upload, file num:0, path:%s", r.URL.Path), wh_)
    }
    f, err := r.MultipartForm.File["file"][0].Open()
    if err != nil {
      e.AssertErr(fmt.Errorf("fail:http.fileheader.open, reason: %s, path:%s", err.Error(), r.URL.Path), wh_)
      return
    }
    stream[0] = f
  } else if cntTyp == "application/download" {
    stream[0] = w
  } else {
    stream[0] = nil
  }

  _path := strings.Split(p, "/")
  if len(_path) != 2 {
    e.AssertErr(errors.New("fail: path, path:"+r.URL.Path), wh_)
  }

  recv, ok := s.Recvs[_path[0]]
  if !ok {
    e.AssertErr(errors.New("fail: no recver, name:"+_path[0]), wh_)
  }

  met, ok := recv.Mets[_path[1]]
  if !ok {
    e.AssertErr(errors.New("fail: no method, name:"+_path[1]), wh_)
  }
  parms := coder.ParseParam(e, met, recv)
  e.Assert()
  e.PrintParams(parms...)
  if stream[0] != nil {
    parms = append(parms, reflect.ValueOf(stream[0]))
  }

  f := met.Method.Func
  // res := f.Call(parms)
  e.SetReV(f.Call(parms))
}
