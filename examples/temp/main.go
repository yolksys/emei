package main

import (
  "github.com/yolksys/emei"
)

func main() {
  emei.RegRecver(_sssss)
  emei.Serve()
  return
  // k := &Arith{}
  // svr.RegRecv(k)
  // svr.Serve()
  // time.Sleep(time.Second * 5)
  // svr.Wait()

  // defer func() {
  //   if false {
  //     return
  //   }
  //   if r := recover(); r != nil {
  //     fmt.Printf("Error: %+v", string(debug.Stack()))
  //   }
  //   //  recover()
  // }()
}

func init() {
 emei.AssertCmd()
}
