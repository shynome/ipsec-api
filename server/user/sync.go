package user

import (
	"net/http"

	"github.com/shynome/ipsec-api/vpn"
	server "github.com/shynome/ipsec-api/server/common"
)

func initSyncHandler() {
	APIMux.HandleFunc("/user/sync", func(w http.ResponseWriter, r *http.Request) {
		err := vpn.Sync()
		server.Resp(w, nil, err)
	})
}
