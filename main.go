package main

import (
	"log"
	"os"

	"github.com/mewfinity06/mewvm/lexer"
	"github.com/mewfinity06/mewvm/runner"
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

	l := lexer.NewLexer(content)

	program, err := l.MakeProgram()
	if err != nil {
		log.Fatal(err)
	}

	r := runner.NewRunner(program)

	_, err = r.Run(true)
	if err != nil {
		log.Fatal(err)
	}
}
