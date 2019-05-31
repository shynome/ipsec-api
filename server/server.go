package server

import (
	"encoding/json"
	"net/http"
)

// APIMux export
var APIMux = http.NewServeMux()

func resp(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	jsondata, err := json.Marshal(data)
	if err != nil {
		resp(w, nil, err)
		return
	}
	headers := w.Header()
	headers["Content-Type"] = []string{"application/json"}
	w.Write(jsondata)
}

func parseParamsFromReq(r *http.Request, v interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&v)
	return
}

func init() {
	initAddHandler()
	initGetpasswordHandler()
	initListHandler()
	initSyncHandler()
}
