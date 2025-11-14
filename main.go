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

	l := lexer.NewLexer(content)
	tokens := make([]lexer.Token, 0)
	for {
		token, err := l.Next()
		if err != nil {
			log.Fatal(err)
		}
		tokens = append(tokens, *token)
		if token.Kind == lexer.TK_Eof {
			break
		}
	}

	for i, token := range tokens {
		fmt.Printf("%d: %v\n", i, token)
	}
}
