package web

import (
  "net/http"

  "github.com/emei/rpc/rpcabs"
)

type webrpc struct {
  // met = http/https
  rpcabs.RpcImp
  real *http.Server
}
