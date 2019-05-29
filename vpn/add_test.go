package vpn

import (
	"testing"
)

func Test_Add(t *testing.T) {
	u := "test3"
	if err := Add(u, "helloworld3"); err != nil {
		t.Error(err)
		return
	}
	exist, err := checkUserExist(u)
	if err != nil {
		t.Error(err)
		return
	}
	if !exist {
		t.Error("用户未添加成功")
		return
	}
	return
}
