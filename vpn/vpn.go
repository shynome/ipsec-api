package vpn

import (
	"os/exec"
	"sync"
)

var mux = sync.Mutex{}

// Add vpn user
func Add(user, password string) error {
	mux.Lock()
	defer mux.Unlock()
	cmd := exec.Command("./adduser", user, password)
	// switch{
	// case <-time.After(5 * time.Second):
	// 	return fmt.Errorf(format string, a ...interface{})
	// }
	cmd.Run()
	return nil
}

// ChangePassword of user
func ChangePassword(user, password string) error {
	return nil
}

// Remove a user
func Remove(user string) error {
	return nil
}

// Sync from ldap server
func Sync(user string) error {
	return nil
}
