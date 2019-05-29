package vpn

import (
	"bufio"
	"os"
	"strings"
)

// List users
// if user not empty, will return the only one user, if not exits return empty
func List(queryUser string) (users []string, err error) {
	inFile, err := os.Open(ipsecConfigFilepath)
	if err != nil {
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {

		userLine := scanner.Text()
		result := strings.SplitN(userLine, ":", 2)
		if result[0] == "" {
			continue
		}
		users = append(users, result[0])

		if queryUser != "" && result[0] == queryUser {
			return
		}

	}

	if queryUser != "" {
		users = []string{}
		return users, nil
	}

	return
}