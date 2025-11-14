package lexer

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrorUnreachable = errors.New("unreachable")

type TokenKind int

// TODO:
const (
	// Keywords
	TK_Add TokenKind = iota
	TK_Push

	// Types
	TK_Int
	TK_Hex

	TK_Eof
)

func stringToTokenKind(s string) (TokenKind, error) {
	switch s {
	case "add":
		return TK_Add, nil
	case "push":
		return TK_Push, nil
	default:
		return TK_Eof, fmt.Errorf("incompatable string found: %s", s)
	}
}

func (tk TokenKind) ToString() string {
	switch tk {
	case TK_Add:
		return "add"
	case TK_Push:
		return "push"
	case TK_Int:
		return "int"
	case TK_Hex:
		return "hex"
	case TK_Eof:
		return "eof"
	}
	log.Fatal(ErrorUnreachable)
	return "unreachable"
}

type Token struct {
	Kind TokenKind
	Word string
}

func (t Token) String() string {
	return fmt.Sprintf("{0x%X => %s}", t.Kind, t.Word)
}

func (t Token) ToHex() (int8, error) {
	switch t.Kind {
	case TK_Add, TK_Push, TK_Eof:
		return int8(t.Kind), nil
	case TK_Hex:
		chopped := strings.TrimPrefix(t.Word, "0x")
		parsed, err := strconv.ParseInt(chopped, 16, 8)
		if err != nil {
			return 0x0, nil
		}
		return int8(parsed), nil
	case TK_Int:
		parsed, err := strconv.ParseInt(t.Word, 10, 8)
		if err != nil {
			return 0x0, err
		}
		return int8(parsed), nil
	default:
		return 0x0, ErrorUnreachable
	}
}

type Lexer struct {
	content []byte
	cur     int
}

func NewLexer(content []byte) Lexer {
	return Lexer{
		content: content,
		cur:     0,
	}
}

func (lex *Lexer) Next() (*Token, error) {
	if lex.cur >= len(lex.content) {
		return &Token{TK_Eof, "eof"}, nil
	}

	start := lex.cur
	c := lex.content[start]
	switch {
	// Skip space
	case unicode.IsSpace(rune(c)):
		lex.cur += 1
		return lex.Next()
	// Read word
	case 'a' <= c && c <= 'z':
		len, err := lex.readIdent()
		if err != nil {
			return nil, err
		}
		word := string(lex.content[start : start+len])
		lex.cur += len

		tk, err := stringToTokenKind(word)
		if err != nil {
			return nil, err
		}

		return &Token{tk, word}, nil

	// Read number
	case '0' <= c && c <= '9':
		len, err := lex.readNum()
		if err != nil {
			return nil, err
		}
		word := string(lex.content[start : start+len])
		lex.cur += len

		if strings.HasPrefix(word, "0x") {
			return &Token{TK_Hex, word}, nil
		} else {
			return &Token{TK_Int, word}, nil
		}
	default:
		return nil, fmt.Errorf("unhandled character: `%c`", c)
	}
}

func (lex *Lexer) readIdent() (int, error) {
	read := 0
	for i := lex.cur; i < len(lex.content); i++ {
		if unicode.IsSpace(rune(lex.content[i])) {
			break
		}
		read += 1
	}
	return read, nil
}

func (lex *Lexer) readNum() (int, error) {
	read := 0
	for i := lex.cur; i < len(lex.content); i++ {
		if unicode.IsSpace(rune(lex.content[i])) {
			break
		}
		read += 1
	}
	return read, nil
}

// Program
type Program []int8

func (l Lexer) MakeProgram() (Program, error) {
	program := make(Program, 0)

	for {
		token, err := l.Next()
		if err != nil {
			return program, err
		}

		hex, err := token.ToHex()
		if err != nil {
			return program, err
		}

		program = append(program, hex)

		if token.Kind == TK_Eof {
			break
		}
	}

	return program, nil
}
