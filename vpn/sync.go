package vpn

// Sync from ldap server
func Sync(queryUser string) (err error) {
	mux.Lock()
	defer mux.Unlock()
	return nil
}
