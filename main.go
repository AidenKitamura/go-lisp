package main

import (
	"bufio"
	"fmt"
	"os"
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
		tokens := GEN_TOKENS(line)
		fmt.Printf("Your Parsed Tokens Are: %v\n", tokens)
		fmt.Printf("Your evaluated results are: %v\n", GEN_EXPRS(tokens).EVAL())
		logger()
	}
}

func logger() {
	fmt.Printf("%s", PREFIX)
}
