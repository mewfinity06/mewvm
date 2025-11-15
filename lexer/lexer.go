package lexer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Errors
var (
	ErrorUnreachable = errors.New("unreachable")
)

type Lexeme interface {
	ToHex() byte
}

// Register
type Register byte

const (
	Reg_A Register = iota
	Reg_B
	Reg_C

	Reg_Math
	Reg_Zero
	Reg_Count // add one
)

func (r Register) ToHex() byte {
	return byte(r)
}

// Instruction
type Instuction byte

const (
	// Math instructions
	Inst_Add Instuction = iota

	// Stack instructions
	Inst_Push // push <operand>
	Inst_Pop  // pop  (reg <r>)
)

func (i Instuction) ToHex() byte {
	return byte(i)
}

// Operand
type Operand byte

const (
	Op_Hex Operand = iota
	Op_Int

	Op_EOF
)

func (o Operand) ToHex() byte {
	return byte(o)
}

// Token type
type Token struct {
	Lexme Lexeme
	Word  string
}

var Token_EOF = Token{
	Lexme: Op_EOF,
	Word:  "eof",
}

type Lexer struct {
	Content [][]byte
	Cur     int
}

func LexerNew(content []byte) Lexer {
	res := Lexer{
		Content: make([][]byte, 0),
		Cur:     0,
	}

	for _, word := range strings.Fields(string(content)) {
		res.Content = append(res.Content, []byte(word))
	}

	return res
}

func (l *Lexer) Next() (*Token, error) {
	if l.Cur >= len(l.Content) {
		return &Token_EOF, nil
	}

	defer func() {
		l.Cur += 1
	}()

	switch word := string(l.Content[l.Cur]); word {
	// Register
	case "reg":
		if l.Cur+1 >= len(l.Content) {
			return &Token_EOF, errors.New("expected register name, found eof")
		}
		l.Cur += 1
		reg := l.Content[l.Cur]
		switch reg := string(reg); reg {
		case "a", "A":
			return &Token{Reg_A, reg}, nil
		case "b", "B":
			return &Token{Reg_B, reg}, nil
		case "c", "C":
			return &Token{Reg_C, reg}, nil
		case "math", "Math":
			return &Token{Reg_Math, reg}, nil
		case "zero", "Zero":
			return &Token{Reg_Zero, reg}, nil
		default:
			return nil, fmt.Errorf("unknown register: %s", reg)
		}
	// Instructions
	case "add":
		return &Token{Inst_Add, word}, nil
	case "push":
		return &Token{Inst_Push, word}, nil
	case "pop":
		return &Token{Inst_Pop, word}, nil
	default:
		// Operands
		// -- Hex
		if strings.HasPrefix(word, "0x") {
			// TODO: Better error checking
			return &Token{Op_Hex, word}, nil
		}

		// -- Int
		if _, err := strconv.ParseInt(word, 10, 0); err == nil {
			return &Token{Op_Int, word}, nil
		}

		return nil, fmt.Errorf("unknown word: %s", word)
	}
}
