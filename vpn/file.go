package vpn

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
)

const fileMode = 0644

var mux = sync.Mutex{}

func appendFile(file, content string) (err error) {

	mux.Lock()
	defer mux.Unlock()

	inFile, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, fileMode)
	if err != nil {
		return
	}
	defer inFile.Close()

	content += "\n"
	if _, err = inFile.WriteString(content); err != nil {
		return
	}

	return

}

func backupFile(file string) (err error) {
	var bakFile, inFile *os.File
	bakFilepath := file + `.bak`

	if inFile, err = os.OpenFile(file, os.O_RDONLY, fileMode); err != nil {
		return
	}
	defer inFile.Close()

	// backup file
	if bakFile, err = os.OpenFile(bakFilepath, os.O_CREATE|os.O_RDWR, fileMode); err != nil {
		return
	}
	defer bakFile.Close()

	if _, err = io.Copy(bakFile, inFile); err != nil {
		return
	}

	return
}

// replaceContent 为空时删除对应的行
func replaceFile(file string, startLinePrefix []string, replaceContent string) (err error) {

	mux.Lock()
	defer mux.Unlock()

	tmpFilepath := file + `.tmp`

	var inFile, tmpFile *os.File
	if inFile, err = os.Open(file); err != nil {
		return
	}
	if tmpFile, err = os.OpenFile(tmpFilepath, os.O_CREATE|os.O_RDWR, fileMode); err != nil {
		return
	}

	checkHasPrefix := func(line string) bool {
		for _, prefix := range startLinePrefix {
			if strings.HasPrefix(line, prefix) {
				return true
			}
		}
		return false
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		if checkHasPrefix(line) {
			if replaceContent == "" {
				continue
			}
			line = replaceContent + "\n"
		} else {
			line += "\n"
		}
		tmpFile.WriteString(line)
	}
	if err = tmpFile.Close(); err != nil {
		return
	}

	if err = inFile.Close(); err != nil {
		return
	}

	if err = backupFile(file); err != nil {
		return
	}

	if err = os.Rename(tmpFilepath, file); err != nil {
		return
	}

	return

}
