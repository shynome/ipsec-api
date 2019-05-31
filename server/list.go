package server

import (
	"net/http"

	"github.com/shynome/ipsec-api/vpn"
)

func initListHandler() {
	APIMux.HandleFunc("/user/list", func(w http.ResponseWriter, r *http.Request) {
		users, err := vpn.List([]string{})
		resp(w, users, err)
	})
}
