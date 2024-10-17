package main

import (
	"errors"
	"strconv"
)

// Parse args to a PasswordArgs struct
func parsePasswordFlags(passwordArgs *PasswordArgs, args ...string) error {
	passLength := DefaultPassLen
	dbPath := DefaultDBPath
	for i, arg := range args {
		if arg == "-s" && i+1 < len(args) {
			passwordArgs.serviceName = args[i+1]
		}

		if arg == "-u" && i+1 < len(args) {
			passwordArgs.userName = args[i+1]
		}

		if arg == "-l" && i+1 < len(args) {
			var err error
			passLength, err = strconv.Atoi(args[i+1])
			if err != nil {
				return err
			}
		}

		if arg == "-db" && i+1 < len(args) {
			dbPath = args[i+1]
		}

		if arg == "-p" && i+1 < len(args) {
			passwordArgs.dbPassword = args[i+1]
		}
	}

	if passwordArgs.dbPassword == "" {
		return errors.New("please provide the db password after the -p flag")
	}

	if passwordArgs.serviceName == "" {
		return errors.New("please provide a service to generate password after the -s flag")
	}

	if passwordArgs.userName == "" {
		return errors.New("please provide a user name to bind the password to for that service, after the -u flag")
	}

	passwordArgs.passLength = passLength
	passwordArgs.dbPath = dbPath
	return nil
}
