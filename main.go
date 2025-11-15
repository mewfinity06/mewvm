package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mewfinity06/mewvm/lexer"
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
	for {
		token, err := l.Next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(token)
		if token.Lexme == lexer.Op_EOF {
			break
		}
	}
}
