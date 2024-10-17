package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
)

type PasswordArgs struct {
	serviceName string
	userName    string
	passLength  int
	dbPath      string
}

// Generates and saves passwod in the db
func generatePassword(args ...string) error {
	passwordArgs := PasswordArgs{}
	err := parseGeneratePasswordFlags(&passwordArgs, args...)

	if err != nil {
		return err
	}

	password := generateRandomString(&passwordArgs.passLength)

	w := tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
	fmt.Fprintln(w, "Service\tUsername\tPassword")
	fmt.Fprintf(w, "%s\t%s\t%s\n", passwordArgs.serviceName, passwordArgs.userName, password)
	w.Flush()

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

// Parse args to a PasswordArgs struct
func parseGeneratePasswordFlags(passwordArgs *PasswordArgs, args ...string) error {
	passLength := DefaultPassLen
	dbPath := DefaultDBPath
	var err error
	for i, arg := range args {
		if arg == "-s" && i+1 < len(args) {
			passwordArgs.serviceName = args[i+1]
		}
		if arg == "-u" && i+1 < len(args) {
			passwordArgs.userName = args[i+1]
		}

		if arg == "-l" && i+1 < len(args) {
			passLength, err = strconv.Atoi(args[i+1])
			if err != nil || passLength < 8 {
				return errors.New("please provide a valid password length greater than 7 after -l flag")
			}
		}

		if arg == "-db" && i+1 < len(args) {
			dbPath = args[i+1]
		}
	}

	if passwordArgs.serviceName == "" {
		return errors.New("please provide a service to generate password after -s flag")
	}

	if passwordArgs.userName == "" {
		return errors.New("please provide a user name to bind the password to for that service after -u flag")
	}
	passwordArgs.passLength = passLength
	passwordArgs.dbPath = dbPath
	return nil
}
