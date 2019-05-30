package vpn

import (
	"testing"
)

func Test_Remove(t *testing.T) {
	var err error
	if err = Remove([]string{"test1"}); err != nil {
		return
	}
	exists, err := checkUserExist("test1")
	if err != nil {
		return
	}
	if exists {
		t.Error("删除用户失败")
		return
	}

	return

}
