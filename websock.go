package emei

import "time"

type WebSock interface {
  ReadMessag() (msgTyp int, b []byte, err error)
  WriteMessag(msgTyp int, b []byte) error
  SetDeadTime(t time.Time) error
  SetReadDeadTime(t time.Time) error
  SetWriteDeadTime(t time.Time) error
  Close() error
}
