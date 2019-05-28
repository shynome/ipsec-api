package vpn

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
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

func timeoutCmd(cmd *exec.Cmd, waitSecond time.Duration) {
	<-time.After(waitSecond)
	if cmd.ProcessState.Exited() {
		return
	}
	cmd.Process.Kill()
}

func encryptoPassword(pass string) (passEnc string, err error) {
	if err = CheckPassword(pass); err != nil {
		return
	}
	cmd := exec.Command("openssl", "passwd", "-1", pass)
	go timeoutCmd(cmd, 5*time.Second)
	passEncBytes, err := cmd.Output()
	if err != nil {
		return
	}
	passEnc = string(passEncBytes)
	return
}

func formatPassword(user, pass, passEnc string) (ppp, ipsec string) {
	ppp = fmt.Sprintf(`"%v" l2tpd "%v." *`, user, pass)
	ipsec = fmt.Sprintf(`%v:%v:xauth-psk`, user, passEnc)
	return
}
