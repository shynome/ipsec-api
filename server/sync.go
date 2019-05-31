package server

import (
	"net/http"

	"github.com/shynome/ipsec-api/vpn"
)

func initSyncHandler() {
	APIMux.HandleFunc("/user/sync", func(w http.ResponseWriter, r *http.Request) {
		err := vpn.Sync()
		resp(w, nil, err)
	})
}
