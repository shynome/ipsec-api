package ldap

import (
	"net/http"
)

// APIMux export
var APIMux = http.NewServeMux()

func init() {
	initListHandler()
}
