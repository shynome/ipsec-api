package common

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
