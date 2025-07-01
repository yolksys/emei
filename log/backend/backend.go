package backend

import (
  _ "github.com/emei/log/backend/console"
  "github.com/emei/log/backend/intra"
  _ "github.com/emei/log/backend/otel"
)

func Get(name string) intra.BcdNew {
  f, _ := intra.RegisteredBck[name]
  return f
}

type Backend = intra.Backend
