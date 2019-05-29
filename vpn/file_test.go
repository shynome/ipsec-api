package vpn

import (
	"testing"
)

const testFilepath = "./test-etc/tmp/test.txt"

func Test_appendFile(t *testing.T) {
	err := appendFile(testFilepath, "hello world")
	err = appendFile(testFilepath, "hello world")
	if err != nil {
		t.Error(err)
		return
	}
	return
}

func Test_replaceFile(t *testing.T) {
	err := replaceFile(testFilepath, `hello world`, "eeee")
	if err != nil {
		t.Error(err)
		return
	}
	return
}
