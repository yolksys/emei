package svr

import (
  "context"
  "reflect"
  "strings"

  "github.com/emei/log"
  "github.com/emei/otel"
  "github.com/emei/rpc"
  "github.com/emei/server/reger"
)

// init(arg ...
func init() {
  // for {reger.RegRpc()}
  for _, value := range rpc.RpcImps {
    reger.RegRpc(value)
  }
}

// RegRcvr ...
func RegRcvr(rcvr any) error {
  r, err := rpc.ParseRcvr(rcvr)
  if err != nil {
    return err
  }
  _rcvrs[strings.ToLower(r.Name)] = r
  return nil
}

// Serve ...
func Serve() error {
  for _, v := range reger.RegedRpc {
    v.RegRcvr(_rcvrs)
  }
  addMetrics()
  for _, v := range reger.RegedRpc {
    v.Start()
    // utils.AssertErr(e)
  }

  wait()
  return nil
}

// wait ...
func wait() {
  l := log.New(context.Background())

  i, v, ok := reflect.Select(reger.SelectRpcErr)
  r := reger.RegedRpc[i]
  l.Error(v)
  if !r.IsClosed() {
    r.Start()
  }

  if !ok {
  }
}

// server recver func
// func (r *Recv)(ev *env.Env, other args...)(rets..., error)

// addMetrics ...
func addMetrics() {
  for _, value := range reger.RegedRpc {
    for _, value0 := range _rcvrs {
      for key1 := range value0.Mets {
        otel.AddApiMeter(value.Name(), value0.Name+"."+key1)
      }
    }
  }
}

var _rcvrs = map[string]*rpc.Recver{}
