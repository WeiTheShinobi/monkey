package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Printf("Hi! This is Monkey programming language REPL")
	repl.Start(os.Stdin, os.Stdout)
}
