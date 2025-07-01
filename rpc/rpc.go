package rpc

import (
  "flag"

  "github.com/emei/rpc/grpc"
  "github.com/emei/rpc/nrpc"
  "github.com/emei/rpc/rpcabs"
  "github.com/emei/rpc/web"
)

type (
  RPC    = rpcabs.RPC
  Recver = rpcabs.Recver
)

func init() {
}

var (
  RpcImps   = map[string]RPC{"web": web.NewRpc(), "nrpc": nrpc.NewRpc(), "grpc": grpc.NewRpc()}
  ParseRcvr = rpcabs.ParseRcvr
  _flagSet  = flag.NewFlagSet("rpc", flag.PanicOnError)
)
