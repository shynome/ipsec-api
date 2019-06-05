package vpn

// SyncUsers will 
type SyncUsers struct {
	Add    []string
	Delete []string
}

// Sync delete user from ldap server
func Sync(confirm bool) (syncUsers SyncUsers, err error) {
	var ldapUsers, ipsecUsers, deleteUsers, addUsers []string
	if ipsecUsers, err = List([]string{}); err != nil {
		return
	}
	if ldapUsers, err = Ldap.GetUsers(); err != nil {
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

	if addUsers, err = findNotExistUsers(ldapUsers); err != nil {
		return
	}

	syncUsers = SyncUsers{
		Add:    addUsers,
		Delete: deleteUsers,
	}

	if confirm == false {
		return
	}

	if len(addUsers) != 0 {
		if err = Add(addUsers); err != nil {
			return
		}
	}
	if len(deleteUsers) != 0 {
		if err = Remove(deleteUsers); err != nil {
			return
		}
	}

	return
}
