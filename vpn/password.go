package vpn

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/uuid"
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
	length := len(passEnc)
	if passEnc[length-2:] == "\r\n" {
		passEnc = passEnc[:length-2]
		return
	}
	if passEnc[length-1:] == "\n" {
		passEnc = passEnc[:length-1]
		return
	}
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

	l2tpdStart, ipsecStart := formatUserPrefix(user)

	replaceFile(l2tpdCoonfigFilepath, []string{l2tpdStart}, l2tpd)
	replaceFile(ipsecConfigFilepath, []string{ipsecStart}, ipsec)

	return
}

// GetPassword of a user
func GetPassword(user string) (password string, err error) {
	inFile, err := os.OpenFile(l2tpdCoonfigFilepath, os.O_RDONLY, fileMode)
	if err != nil {
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	user = fmt.Sprintf(`"%v"`, user)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, user) {
			continue
		}
		result := strings.SplitN(line, `l2tpd "`, 2)
		result = strings.SplitN(result[1], `"`, 2)
		pass := result[0]
		if len(pass) < 1 {
			continue
		}
		password = pass
		return
	}

	err = fmt.Errorf("can't found the user %v password", user[1:len(user)-1])

	return
}

// GetShareSecret of ipsec
func GetShareSecret() (shareSecret string, err error) {
	inFile, err := os.OpenFile(ipsecSecretsFilepath, os.O_RDONLY, fileMode)
	if err != nil {
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		result := strings.SplitN(line, `PSK "`, 2)
		secret := result[1]
		if len(secret) < 2 {
			continue
		}
		shareSecret = secret[0 : len(secret)-1]
		return
	}
	err = fmt.Errorf("can't found ipsec secret")
	return
}

// GenPassword generate a password
func GenPassword() string {
	return uuid.New().String()
}
