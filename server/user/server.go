package user

import (
	"net/http"
)

// APIMux export
var APIMux = http.NewServeMux()

func init() {
	initAddHandler()
	initGetpasswordHandler()
	initListHandler()
	initSyncHandler()
}
