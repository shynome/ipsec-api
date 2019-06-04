package user

import (
	"fmt"
	"net/http"

	"github.com/shynome/ipsec-api/vpn"
	server "github.com/shynome/ipsec-api/server/common"
)

// GetpasswordHandlerParams params
type GetpasswordHandlerParams struct {
	User string
}

func initGetpasswordHandler() {
	APIMux.HandleFunc("/user/getpassword", func(w http.ResponseWriter, r *http.Request) {
		var err error
		params := &GetpasswordHandlerParams{}
		if err = server.ParseParamsFromReq(r, params); err != nil {
			server.Resp(w, nil, err)
			return
		}
		if params.User == "" {
			err := fmt.Errorf("required user value")
			server.Resp(w, nil, err)
			return
		}
		password, err := vpn.GetPassword(params.User)
		if err != nil {
			server.Resp(w, nil, err)
			return
		}
		server.Resp(w, map[string]string{"password": password}, nil)
		return
	})
}
