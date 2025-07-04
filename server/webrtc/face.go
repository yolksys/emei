package webrtc

// signal class from 1------------------------------
const (
  // req attrs: jwt
  // optional retv attrs:  code, reason
  SigClassLogin byte = iota + 1
  // req attrs: name=uuid target=user name
  // optional retv attrs: code, reason
  SigClassOTOCaller // one to one
  // req attrs: name=uuid target=uuid
  // optional retv attrs: code, reason
  SigClassOTOCallee // one to one
  // req attrs: passd,  name=uuid
  // optional retv attrs: code, reason
  SigClassAllocOTM // one to multiple
  SigClassAllocMTM // multi to multi
  // Lecture:Students can only see the teacher
  // Or can see vidio and hear voice from the student who is currently speaking
  // teachers can see all students
  SigClassAllocLCT
  // req attrs: passd, name, target="target peerid"
  // optional retv attrs: code, reason
  SigClassAllocCli = 15 // Participants
)

// req attrs: name="self uuid"
const (
  SigClassCfg byte = iota + 16
)

const (
  SigClassNotify   byte = 250
  SigClassResponse byte = 255
)

// signal type from 1---------------------------------
const (
  SigTypeStmWebrtc byte = iota + 1
  SigTypeStmRTP
  SigTypeStmRTMP

  SigStmTypeCustomTCP  = 13
  SigStmTypeCustomUDP  = 14
  SigStmTypeCustomQUIC = 15
)

const (
  SigTypeCfgSDP byte = iota + 16
  SigTypeCfgCadidate
)

// attrs: caller="name"
const SigTypeNtyCaller byte = iota + 235
