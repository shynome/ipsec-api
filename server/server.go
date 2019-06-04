package server

import (
	"net/http"

	"github.com/shynome/ipsec-api/server/ldap"
	"github.com/shynome/ipsec-api/server/user"
)

// APIMux export
var APIMux = http.NewServeMux()

func init() {
	APIMux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		user.APIMux.ServeHTTP(w, r)
	})
	APIMux.HandleFunc("/ldap/", func(w http.ResponseWriter, r *http.Request) {
		ldap.APIMux.ServeHTTP(w, r)
	})
}
