package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mewfinity06/mewvm/lexer"
	"github.com/mewfinity06/mewvm/packer"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Not enough args")
	}

	content, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	l := lexer.LexerNew(content)

	program, err := packer.Pack(l)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Program: %v\n", program)
}
