package user

import (
	"net/http"

	server "github.com/shynome/ipsec-api/server/common"
	"github.com/shynome/ipsec-api/vpn"
)

func initListHandler() {
	APIMux.HandleFunc("/user/list", func(w http.ResponseWriter, r *http.Request) {
		users, err := vpn.List([]string{})
		server.Resp(w, users, err)
	})
}
