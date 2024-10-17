package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func displayPassword(passwordArgs *PasswordArgs, password string) {
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
	fmt.Fprintln(w, "Service\tUsername\tPassword")
	fmt.Fprintf(w, "%s\t%s\t%s\n", passwordArgs.serviceName, passwordArgs.userName, password)
	w.Flush()
}
