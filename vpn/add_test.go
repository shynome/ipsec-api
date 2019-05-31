package vpn

import (
	"testing"
)

func Test_Add(t *testing.T) {
	users := []string{"test3", "test4"}
	if err := Add(users); err != nil {
		t.Error(err)
		return
	}
	notExistUsers, err := findNotExistUsers(users)
	if err != nil {
		t.Error(err)
		return
	}
	if len(notExistUsers) != 0 {
		t.Error("用户未添加成功")
		return
	}
	return
}
