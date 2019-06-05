package user

import (
	"net/http"

	server "github.com/shynome/ipsec-api/server/common"
	"github.com/shynome/ipsec-api/vpn"
)

// SyncHandlerParams params
type SyncHandlerParams struct {
	Confirm bool
}

func initSyncHandler() {
	APIMux.HandleFunc("/user/sync", func(w http.ResponseWriter, r *http.Request) {
		var err error
		params := &SyncHandlerParams{}
		if err = server.ParseParamsFromReq(r, params); err != nil {
			server.Resp(w, nil, err)
			return
		}
		syncUsers, err := vpn.Sync(params.Confirm)
		server.Resp(w, syncUsers, err)
	})
}
