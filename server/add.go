package server

import (
	"fmt"
	"net/http"

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
		if err = parseParamsFromReq(r, params); err != nil {
			resp(w, nil, err)
			return
		}
		if len(params.Users) == 0 {
			err = fmt.Errorf("can't add empty user list")
			resp(w, nil, err)
			return
		}
		if err = vpn.Add(params.Users); err != nil {
			resp(w, nil, err)
			return
		}
		resp(w, nil, nil)
	})
}
