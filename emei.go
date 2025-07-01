package emei

import (
  "time"

  "github.com/emei/cfg"
  "github.com/emei/cmd"
  "github.com/emei/env"
  "github.com/emei/log"
  _c "github.com/emei/rpc/call"
  svr "github.com/emei/server"
)

// env
type Env interface {
  Return()
  Assert()
  HasError() bool
  Err() error
  ResetErr()
  AssertErr(err error, clear ...func())
  Errorf(code uint16, f string, args ...any) error
  Event(args ...interface{})
}

var NewEnv = env.New

// server
var (
  RegRecver = svr.RegRcvr
  Serve     = svr.Serve
)

// caller
var (
  CallWithRStream  = _c.CallWithRStream
  CallWithWStream  = _c.CallWithWStream
  CallWithRWStream = _c.CallWithRWStream
)

func call(e Env, svcName, met string, args ...any) error {
  return _c.Call(e.(env.Env), svcName, met, args...)
}

func Call1[T1 any](e Env, svcName, met string, args ...any) (T1, error) {
  return _c.Call1[T1](e.(env.Env), svcName, met, args...)
}

func Call2[T1, T2 any](e Env, svcName, met string, args ...any) (T1, T2, error) {
  return _c.Call2[T1, T2](e.(env.Env), svcName, met, args...)
}

func Call3[T1, T2, T3 any](e Env, svcName, met string, args ...any) (T1, T2, T3, error) {
  return _c.Call3[T1, T2, T3](e.(env.Env), svcName, met, args...)
}

// cmd
var (
  AssertCmd = cmd.Assert
  BoolCmd   = cmd.Bool
  StringCmd = cmd.String
)

// config
var (
  GetCfgItem = cfg.GetCfgItem
  NewCfg     = cfg.New
)

// log
var (
  Log    = log.Log
  NewLog = log.New
)

// websock
type WebSock interface {
  ReadMessag() (msgTyp int, b []byte, err error)
  WriteMessag(msgTyp int, b []byte) error
  SetDeadTime(t time.Time) error
  SetReadDeadTime(t time.Time) error
  SetWriteDeadTime(t time.Time) error
  Close() error
}
