package ldap

import (
	"net/http"

	server "github.com/shynome/ipsec-api/server/common"
	"github.com/shynome/ipsec-api/vpn"
)

func initListHandler() {
	APIMux.HandleFunc("/ldap/list", func(w http.ResponseWriter, r *http.Request) {
		users, err := vpn.Ldap.GetUsers()
		server.Resp(w, users, err)
	})
}
