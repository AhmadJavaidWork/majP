package main

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/ahmadjavaidwork/majP/internal/auth"
	"github.com/ahmadjavaidwork/majP/internal/encrypt"
)

// Returns comma separated string of service, username and password if the entry exists
// otherwise returns empty string
func getPasswordEntry(passwordArgs *PasswordArgs) (string, error) {
	if _, err := os.Stat("pass_hash"); errors.Is(err, os.ErrNotExist) {
		return "", os.ErrNotExist
	}

	f, err := os.OpenFile("pass_hash", os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	err = auth.CheckPasswordHash(scanner.Text(), passwordArgs.dbPassword)
	if err != nil {
		return "", errors.New("wrong db password")
	}

	f, err = os.OpenFile(passwordArgs.dbPath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		curEntry := strings.Split(encrypt.Decrypt(scanner.Text(), passwordArgs.dbPassword), ",")
		if curEntry[0] == passwordArgs.serviceName && curEntry[1] == passwordArgs.userName {
			return curEntry[2], nil
		}
	}

	return "", nil
}
