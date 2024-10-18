package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/ahmadjavaidwork/majP/internal/auth"
	"github.com/ahmadjavaidwork/majP/internal/encrypt"
)

type PasswordArgs struct {
	serviceName string
	userName    string
	passLength  int
	dbPath      string
	dbPassword  string
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
		if errors.Is(err, os.ErrNotExist) {
			err = createDBPasswordHash(&passwordArgs)
			if err != nil {
				return err
			}
		} else {
			return err
		}
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
	line = encrypt.Encrypt(line, passwordArgs.dbPassword)

	_, err = f.WriteString(line)

	if err != nil {
		return fmt.Errorf("error saving password: %w", err)
	}
	return nil
}

func createDBPasswordHash(passwordArgs *PasswordArgs) error {
	f, err := os.OpenFile("pass_hash", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	defer f.Close()
	passwordHash, err := auth.HashPassword(passwordArgs.dbPassword)
	if err != nil {
		return err
	}
	_, err = f.WriteString(passwordHash)

	if err != nil {
		return fmt.Errorf("error saving password hash: %w", err)
	}
	return nil
}
