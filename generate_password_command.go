package main

import "errors"

func generatePassword(args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide a service to generate password")
	}
	return nil
}
