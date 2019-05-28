package common

import (
	"fmt"
	"strings"
)

// NotAllowedCharacters for password set
const NotAllowedCharacters = `\"'`

// CheckPassword of user
func CheckPassword(pass string) (err error) {
	if strings.IndexAny(pass, NotAllowedCharacters) != -1 {
		return fmt.Errorf("VPN credentials must not contain any of these characters: %v", pass)
	}
	return
}
