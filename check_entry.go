package main

import (
	"bufio"
	"os"
	"strings"
)

// Returns comma separated string of service, username and password if the entry exists
// else returns empty string
func getPasswordEntry(passwordArgs *PasswordArgs) (string, error) {
	f, err := os.OpenFile(passwordArgs.dbPath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curEntry := strings.Split(scanner.Text(), ",")
		if curEntry[0] == passwordArgs.serviceName && curEntry[1] == passwordArgs.userName {
			return curEntry[2], nil
		}
	}

	return "", nil
}
