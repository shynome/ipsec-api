package vpn

import (
	"fmt"
)

func checkUserExist(user string) (exist bool, err error) {
	exist = true
	users, err := List([]string{user})
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

func findNotExistUsers(users []string) (notExistUsers []string, err error) {

	existUsers, err := List(users)
	if err != nil {
		return
	}
	existUserMaps := map[string]bool{}
	for _, user := range existUsers {
		existUserMaps[user] = true
	}

	for _, user := range users {
		if existUserMaps[user] {
			continue
		}
		notExistUsers = append(notExistUsers, user)
	}

	return
}

// Add vpn user
func Add(users []string) (err error) {

	notExistUsers, err := findNotExistUsers(users)
	if err != nil {
		return
	}
	if len(notExistUsers) == 0 {
		err = fmt.Errorf("all users is exist")
		return
	}

	var l2tpdUsers, ipsecUsers string
	for _, user := range notExistUsers {
		pass := GenPassword()
		var passEnc string
		passEnc, err = encryptoPassword(pass)
		if err != nil {
			return
		}
		l2tpd, ipsec := formatUser(user, pass, passEnc)
		l2tpdUsers += l2tpd + "\n"
		ipsecUsers += ipsec + "\n"
	}

	appendFile(l2tpdCoonfigFilepath, l2tpdUsers[:len(l2tpdUsers)-1])
	appendFile(ipsecConfigFilepath, ipsecUsers[:len(ipsecUsers)-1])

	return nil
}
