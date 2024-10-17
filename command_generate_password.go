package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type PasswordArgs struct {
	serviceName string
	userName    string
	passLength  int
	dbPath      string
}

// Generates and saves password in the db
func commandGeneratePassword(args ...string) error {
	passwordArgs := PasswordArgs{}
	err := parsePasswordFlags(&passwordArgs, args...)

	if err != nil {
		return err
	}

	if passwordArgs.passLength < 8 || passwordArgs.passLength > 20 {
		return errors.New("please provide a valid password length between 8 and 20")
	}

	passwordEntry, err := getPasswordEntry(&passwordArgs)
	if err != nil {
		return err
	}

	if passwordEntry != "" {
		return errors.New("password already exists")
	}

	password := generateRandomString(&passwordArgs.passLength)
	displayPassword(&passwordArgs, password)
	err = savePassword(&passwordArgs, password)
	if err != nil {
		return err
	}

	return nil
}

// Generate a random string of given length ascii values of range [33, 126]
func generateRandomString(l *int) string {
	s := ""
	for *l > 0 {
		n := rand.Int31n(94) + 33
		s += string(rune(n))
		*l--
	}
	return s
}

// Saves comma separated password service name and username in the provided db path
func savePassword(passwordArgs *PasswordArgs, password string) error {

	f, err := os.OpenFile(passwordArgs.dbPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	defer f.Close()

	line := fmt.Sprintf("%s,%s,%s\n", passwordArgs.serviceName, passwordArgs.userName, password)
	_, err = f.WriteString(line)

	if err != nil {
		return fmt.Errorf("error saving password: %w", err)
	}
	return nil
}
