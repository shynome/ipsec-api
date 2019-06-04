package server

import (
	"net/http"

	"github.com/shynome/ipsec-api/server/ldap"
	"github.com/shynome/ipsec-api/server/user"
)

// APIMux export
var APIMux = http.NewServeMux()

func init() {
	APIMux.Handle("/user/", user.APIMux)
	APIMux.Handle("/ldap/", ldap.APIMux)
}
