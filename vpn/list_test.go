package vpn

import (
	"testing"
)

func TestList(t *testing.T) {
	cases := map[string]int{
		"":      2,
		"test1": 1,
		"test3": 0,
	}
	for queryUser, expectUesrsLength := range cases {
		users, err := List([]string{queryUser})
		if err != nil {
			t.Error(err)
		}
		if len(users) != expectUesrsLength {
			t.Errorf("users length except %v, get %v", expectUesrsLength, len(users))
			t.Error(users)
		}
	}
	return
}
