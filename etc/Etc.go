package etc

import (
  "github.com/emei/etc/dns"
  "github.com/emei/etc/etcintra"
  "github.com/emei/etc/etcout"
)

// Get ...
func GetSvcCfg(svc string) (string, error) {
  t, err := etcintra.EtcCli.Get("service:cfg", "", svc)
  v, _ := t[svc]
  return v, err
}

// Watch ...
func WatchSvcCfg(svc string) (<-chan etcout.Event, error) {
  return etcintra.EtcCli.Watch("service:cfg", "", svc)
}

// GetSvcIp ...
func GetSvcIp(svc string) (string, error) {
  return dns.GetSvcIp(svc)
}
