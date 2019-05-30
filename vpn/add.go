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
		exist = false
		return
	}
	if usersLen != 1 {
		err = fmt.Errorf("user length except 1 , but get %v", usersLen)
		return
	}
	exist = usersLen == 1
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

	passEnc, err := encryptoPassword(pass)
	if err != nil {
		return
	}

	l2tpd, ipsec := formatUser(user, pass, passEnc)

	appendFile(l2tpdCoonfigFilepath, l2tpd)
	appendFile(ipsecConfigFilepath, ipsec)

	return nil
}
