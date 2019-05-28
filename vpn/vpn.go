package vpn

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var mux = sync.Mutex{}

var (
	l2tpdCoonfigFilepath string
	ipsecConfigFilepath  string
)

func init() {
	l2tpdCoonfigFilepath = "/etc/ppp/chap-secrets"
	ipsecConfigFilepath = "/etc/ipsec.d/passwd"
}

// List users
// if user not empty, will return the only one user, if not exits return empty
func List(queryUser string) (users []string, err error) {
	inFile, err := os.Open(ipsecConfigFilepath)
	if err != nil {
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {

		userLine := scanner.Text()
		result := strings.SplitN(userLine, ":", 2)
		if result[0] == "" {
			continue
		}
		users = append(users, result[0])

		if queryUser != "" && result[0] == queryUser {
			return
		}

	}

	if queryUser != "" {
		users = []string{}
		return users, nil
	}

	return
}

// Add vpn user
func Add(user, pass string) (err error) {
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

// Remove a user
func Remove(queryUser string) (err error) {
	mux.Lock()
	defer mux.Unlock()
	return nil
}

// Sync from ldap server
func Sync(queryUser string) (err error) {
	mux.Lock()
	defer mux.Unlock()
	return nil
}
