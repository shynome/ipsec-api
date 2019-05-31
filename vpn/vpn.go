package vpn

import (
	"os"
)

var (
	l2tpdCoonfigFilepath = "/etc/ppp/chap-secrets"
	ipsecConfigFilepath  = "/etc/ipsec.d/passwd"
	ipsecSecretsFilepath = "/etc/ipsec.secrets"
)

func init() {
	if f := os.Getenv("l2tpdCoonfigFilepath"); f != "" {
		l2tpdCoonfigFilepath = f
	}
	if f := os.Getenv("ipsecConfigFilepath"); f != "" {
		ipsecConfigFilepath = f
	}
	if f := os.Getenv("ipsecSecretsFilepath"); f != "" {
		ipsecSecretsFilepath = f
	}
}
