package server

import (
	"net/http"

	"encoding/json"

	"github.com/shynome/ipsec-api/vpn"
)

// APIMux export
var APIMux = http.NewServeMux()

func resp(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if data == nil {

	}
	headers := w.Header()
	headers["Content-Type"] = []string{"application/json"}
}

func init() {
	APIMux.HandleFunc("/user/add", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("add"))
	})
	APIMux.HandleFunc("/user/getpassword", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("getpassword"))
	})
	APIMux.HandleFunc("/user/list", func(w http.ResponseWriter, r *http.Request) {
		users, err := vpn.List("")
		if err != nil {
			resp(w, nil, err)
			return
		}
		res, err := json.Marshal(users)
		if err != nil {
			resp(w, nil, err)
			return
		}
		w.Write(res)
	})
	APIMux.HandleFunc("/user/sync", func(w http.ResponseWriter, r *http.Request) {
		err := vpn.Sync()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
