package web

import "github.com/emei/rpc/call"

func init() {
  call.RegSender("web", newWebSender())
}
