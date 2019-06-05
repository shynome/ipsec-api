package vpn

import (
	"testing"
)

func Test_Sync(t *testing.T) {
	var err error
	var users SyncUsers
	if users, err = Sync(false); err != nil {
		return
	}
	t.Log(users)
	return
}
