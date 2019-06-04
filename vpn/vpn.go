package vpn

import (
	"os"

	"github.com/shynome/ipsec-api/ldap"
)

var (
	l2tpdCoonfigFilepath = "/etc/ppp/chap-secrets"
	ipsecConfigFilepath  = "/etc/ipsec.d/passwd"
	ipsecSecretsFilepath = "/etc/ipsec.secrets"
)

// Ldap instance for vpn
var Ldap = &ldap.LDAP{}

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
	if err := ldap.NewLDAP(Ldap); err != nil {
		panic(err)
	}
}
