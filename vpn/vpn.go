package vpn

import (
	"sync"
)

var mux = sync.Mutex{}

var (
	l2tpdCoonfigFilepath string
	ipsecConfigFilepath  string
)

func init() {
	l2tpdCoonfigFilepath = "/etc/ppp/chap-secrets"
	ipsecConfigFilepath = "/etc/ipsec.d/passwd"
}
