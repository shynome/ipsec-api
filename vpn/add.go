package vpn

import (
	"fmt"
)

func checkUserExist(user string) (exist bool, err error) {
	exist = true
	users, err := List(user)
	if err != nil {
		return
	}
	usersLen := len(users)
	if usersLen == 0 {
		return
	}
	if usersLen != 1 {
		err = fmt.Errorf("user length except 1 , but get %v", usersLen)
		return
	}
	exist = false
	return
}

// Add vpn user
func Add(user, pass string) (err error) {

	exist, err := checkUserExist(user)
	if err != nil {
		return
	}
	if exist {
		err = fmt.Errorf("user %v has exist", user)
		return
	}

	mux.Lock()
	defer mux.Unlock()

	passEnc, err := encryptoPassword(pass)
	if err != nil {
		return
	}

	ppp, ipsec := formatPassword(user, pass, passEnc)

	fmt.Print(ppp, ipsec)

	return nil
}

// ChangePassword of user
func ChangePassword(user, pass string) (err error) {
	mux.Lock()
	defer mux.Unlock()
	return nil
}
