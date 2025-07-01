package grpc

import "github.com/emei/rpc/call"

func init() {
  call.RegSender("grpc", newSender())
}
