package vpn

import (
	"bufio"
	"os"
	"strings"
)

// List users
// if user not empty, will return the only one user, if not exits return empty
func List(queryUser []string) (users []string, err error) {
	inFile, err := os.OpenFile(ipsecConfigFilepath, os.O_RDONLY, fileMode)
	if err != nil {
		return
	}
	defer inFile.Close()

	execQuery := false
	queryUserMaps := map[string]bool{}
	for _, user := range queryUser {
		if user == "" {
			continue
		}
		queryUserMaps[user] = true
		execQuery = true
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {

		userLine := scanner.Text()
		result := strings.SplitN(userLine, ":", 2)
		if result[0] == "" {
			continue
		}

		user := result[0]

		if execQuery {
			findedUser := result[0]
			if queryUserMaps[findedUser] == false {
				continue
			}
			// 完成了的话就提前退出
			if len(users) == len(queryUser) {
				return
			}
		}

		users = append(users, user)

	}

	return
}
