package emei

import "github.com/yolksys/emei/env"

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
