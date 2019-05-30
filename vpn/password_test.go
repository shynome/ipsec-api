package vpn

import (
	"testing"
)

func TestCheckPassword(t *testing.T) {
	cases := map[string]bool{
		`eqrwqrq"`:      false,
		`eqrwqrw323`:    true,
		`ffwefqfqw`:     true,
		`eqrwqrw3ww'`:   false,
		`eqrw"qrw3ww'\`: false,
	}
	for pass, exceptResult := range cases {
		err := CheckPassword(pass)
		checkResult := err == nil
		if checkResult != exceptResult {
			t.Errorf("the password `%v` check result is %v , except result is %v", pass, checkResult, exceptResult)
		}
	}
}

func Test_ChangePassword(t *testing.T) {
	var err error
	if err = ChangePassword("test1", "8888888888"); err != nil {
		t.Error(err)
		return
	}
	return
}

func Test_GetShareSecret(t *testing.T) {
	shareSecret, err := GetShareSecret()
	if err != nil {
		t.Error(err)
		return
	}
	var rightShareSecret = "fewfqerqrwqrqrwqrqwrwqr"
	if shareSecret != rightShareSecret {
		t.Errorf("can't get the right share secret, now get %v", shareSecret)
		return
	}
	return
}

func Test_GetPassword(t *testing.T) {
	password, err := GetPassword("test2")
	if err != nil {
		t.Error(err)
		return
	}
	var rightPassword = "VPN_PASSWORD2"
	if password != rightPassword {
		t.Errorf("can't get the right share secret, now get %v", password)
		return
	}
	return
}
