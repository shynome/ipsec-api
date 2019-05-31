package server

import (
	"fmt"
	"net/http"

	"github.com/shynome/ipsec-api/vpn"
)

// GetpasswordHandlerParams params
type GetpasswordHandlerParams struct {
	user string
}

func initGetpasswordHandler() {
	APIMux.HandleFunc("/user/getpassword", func(w http.ResponseWriter, r *http.Request) {
		var err error
		params := &GetpasswordHandlerParams{}
		if err = parseParamsFromReq(r, params); err != nil {
			resp(w, nil, err)
			return
		}
		if params.user == "" {
			err := fmt.Errorf("required user value")
			resp(w, nil, err)
			return
		}
		password, err := vpn.GetPassword(params.user)
		if err != nil {
			resp(w, nil, err)
			return
		}
		resp(w, map[string]string{"password": password}, nil)
		return
	})
}
