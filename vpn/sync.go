package vpn

import (
	"github.com/shynome/ipsec-api/ldap"
)

// Sync delete user from ldap server
func Sync() (err error) {
	ld := &ldap.LDAP{}
	if err = ldap.NewLDAP(ld); err != nil {
		return
	}
	var ldapUsers, ipsecUsers, deleteUsers []string
	if ipsecUsers, err = List(""); err != nil {
		return
	}
	if ldapUsers, err = ld.GetUsers(); err != nil {
		return
	}
	ldapUserMaps := map[string]interface{}{}
	for _, user := range ldapUsers {
		ldapUserMaps[user] = true
	}

	for _, user := range ipsecUsers {
		if ldapUserMaps[user] == nil {
			deleteUsers = append(deleteUsers, user)
		}
	}

	if err = Remove(deleteUsers); err != nil {
		return
	}

	return
}
