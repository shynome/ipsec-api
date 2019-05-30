package vpn

// Remove a user
func Remove(deleteUsers []string) (err error) {
	l2tpdDeleteUsers, ipsecDeleteUsers := []string{}, []string{}
	for _, user := range deleteUsers {
		l2tpdStart, ipsecStart := formatUserPrefix(user)
		l2tpdDeleteUsers = append(l2tpdDeleteUsers, l2tpdStart)
		ipsecDeleteUsers = append(ipsecDeleteUsers, ipsecStart)
	}

	if err = replaceFile(l2tpdCoonfigFilepath, l2tpdDeleteUsers, ""); err != nil {
		return
	}
	if err = replaceFile(ipsecConfigFilepath, ipsecDeleteUsers, ""); err != nil {
		return
	}
	return
}
