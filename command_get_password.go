package main

import (
	"errors"
	"fmt"
	"os"
)

// Displays password if it exists
func commandGetPassword(args ...string) error {
	passwordArgs := PasswordArgs{}
	err := parsePasswordFlags(&passwordArgs, args...)

	if err != nil {
		return err
	}

	entry, err := getPasswordEntry(&passwordArgs)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("password hash file does not exist")
		}
		return err
	}

	if entry == "" {
		return fmt.Errorf("no password exists for %s and %s", passwordArgs.serviceName, passwordArgs.userName)
	}

	displayPassword(&passwordArgs, entry)
	return nil
}
