package auth

import (
	"io/ioutil"
	"strings"
)

func parseAccountsFile(filePath string) (map[string]string, error) {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	// Read file row by row
	userRows := strings.Split(string(fileContent), "\n")

	// Result structure will be [email] = password
	users := make(map[string]string, len(userRows))

	for _, row := range userRows {

		// Make sure we ignore any trailing/leading whitespace
		row = strings.TrimSpace(row)

		// Ignore comma-rows
		if strings.Index(row, "#") == 0 {
			continue
		}

		user := strings.Split(row, ":")

		// check that the user row is a valid user
		if len(user) != 2 {
			continue
		}

		// Add to the map
		users[user[0]] = user[1]
	}

	return users, nil
}
