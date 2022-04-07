package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AidenKitamura/go-lisp/parser"
)

const (
	PREFIX = "> "
)

func main() {
	fmt.Println("Welcome to Go Lisp Interpreter!")
	scanner := bufio.NewScanner(os.Stdin)
	logger()
	for scanner.Scan() {
		// Do something with the data
		line := scanner.Text()
		tokens := parser.Tokenize(line)
		fmt.Printf("Your Parsed Tokens Are: %v\n", tokens)
		logger()
	}
}

func logger() {
	fmt.Printf("%s", PREFIX)
}
