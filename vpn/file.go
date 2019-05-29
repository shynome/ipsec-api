package vpn

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const fileMode = 0600

func appendFile(file, content string) (err error) {

	inFile, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND, fileMode)
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
func replaceFile(file, startLineContent, replaceContent string) (err error) {

	tmpFilepath := file + `.tmp`

	var inFile, tmpFile *os.File
	if inFile, err = os.Open(file); err != nil {
		return
	}
	if tmpFile, err = os.OpenFile(tmpFilepath, os.O_CREATE|os.O_RDWR, fileMode); err != nil {
		return
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, startLineContent) {
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
