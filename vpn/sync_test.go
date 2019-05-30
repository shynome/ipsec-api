package vpn

import (
	"testing"
)

func Test_Sync(t *testing.T) {
	var err error
	if err = Sync(); err != nil {
		return
	}
	return
}
