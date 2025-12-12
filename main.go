package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/maybe-joe/monkey/repl"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")

	if err := repl.Run(os.Stdin, os.Stdout); err != nil {
		return err
	}

	return nil
}
