package user

import (
	"fmt"
	"net/http"

	server "github.com/shynome/ipsec-api/server/common"
	"github.com/shynome/ipsec-api/vpn"
)

// AddHandlerParams params
type AddHandlerParams struct {
	Users []string
}

func initAddHandler() {
	APIMux.HandleFunc("/user/add", func(w http.ResponseWriter, r *http.Request) {
		var err error
		params := &AddHandlerParams{}
		if err = server.ParseParamsFromReq(r, params); err != nil {
			server.Resp(w, nil, err)
			return
		}
		if len(params.Users) == 0 {
			err = fmt.Errorf("can't add empty user list")
			server.Resp(w, nil, err)
			return
		}
		if err = vpn.Add(params.Users); err != nil {
			server.Resp(w, nil, err)
			return
		}
		server.Resp(w, nil, nil)
	})
}
