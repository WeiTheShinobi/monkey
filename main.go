package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s! This is Monkey programming language REPL", u.Name)
	repl.Start(os.Stdin, os.Stdout)
}
