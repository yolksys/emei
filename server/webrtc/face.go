package webrtc

// class
const (
  SigClassLogin    byte = iota + 1
  SigClassAllocOTO      // one to one
  SigClassAllocOTM      // one to multiple
  SigClassAllocMTM      // multi to multi
  SigClassAllocCli      // client, do not upload
)
const SigClassSwitchCfg byte = iota + 16

// stm type
const (
  SigTypeStmWebrtc byte = iota + 1
  SigTypeStmRTP
  SigTypeStmRTMP
)

const (
  SigTypeCfgSDP byte = iota + 16
  SigTypeCfgCadidate
)
const SigStmTypeCustom = 255
