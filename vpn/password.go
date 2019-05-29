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

func formatUser(user, pass, passEnc string) (l2tpd, ipsec string) {
	l2tpd = fmt.Sprintf(`"%v" l2tpd "%v" *`, user, pass)
	ipsec = fmt.Sprintf(`%v:%v:xauth-psk`, user, passEnc)
	return
}

func formatUserPrefix(user string) (l2tpd, ipsec string) {
	l2tpd = fmt.Sprintf(`"%v" l2tpd`, user)
	ipsec = fmt.Sprintf(`%v:`, user)
	return
}

// ChangePassword of user
func ChangePassword(user, pass string) (err error) {

	exist, err := checkUserExist(user)
	if err != nil {
		return
	}
	if !exist {
		return fmt.Errorf("用户 %v 不存在", user)
	}

	passEnc, err := encryptoPassword(pass)
	if err != nil {
		return
	}

	l2tpd, ipsec := formatUser(user, pass, passEnc)

	mux.Lock()
	defer mux.Unlock()

	l2tpdStart, ipsecStart := formatUserPrefix(user)

	replaceFile(l2tpdCoonfigFilepath, l2tpdStart, l2tpd)
	replaceFile(ipsecConfigFilepath, ipsecStart, ipsec)

	return
}
