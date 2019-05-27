package ldap

import (
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func TestLdap(t *testing.T) {
	ld := &LDAP{}
	err := NewLDAP(ld)
	if err != nil {
		t.Error(err)
		return
	}
	return
}
