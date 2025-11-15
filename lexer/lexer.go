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
	Reg_Count // NOT FOR USE
)

func (r Register) ToHex() byte {
	return byte(r)
}

// Instruction
type Instruction byte

const (
	Inst_Nop Instruction = iota
	Inst_Halt

	// Math instructions
	Inst_Add

	// Stack instructions
	Inst_Push // push <operand>
	Inst_Pop  // pop  (reg <r>)

	Inst_EOF
)

func (i Instruction) ToHex() byte {
	return byte(i)
}

// Operand
type Operand byte

const (
	Op_Hex Operand = iota
	Op_Int
)

func (o Operand) ToHex() byte {
	return byte(o)
}

// Token type
type Token struct {
	Lexme Lexeme
	Word  string
	Value byte
}

var Token_EOF = Token{
	Lexme: Inst_EOF,
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
			return &Token{Reg_A, reg, 0}, nil
		case "b", "B":
			return &Token{Reg_B, reg, 0}, nil
		case "c", "C":
			return &Token{Reg_C, reg, 0}, nil
		case "math", "Math":
			return &Token{Reg_Math, reg, 0}, nil
		case "zero", "Zero":
			return &Token{Reg_Zero, reg, 0}, nil
		default:
			return nil, fmt.Errorf("unknown register: %s", reg)
		}
	// Instructions
	case "nop":
		return &Token{Inst_Nop, word, 0}, nil
	case "halt":
		return &Token{Inst_Halt, word, 0}, nil
	case "add":
		return &Token{Inst_Add, word, 0}, nil
	case "push":
		return &Token{Inst_Push, word, 0}, nil
	case "pop":
		return &Token{Inst_Pop, word, 0}, nil
	default:
		// Operands
		// -- Hex
		if strings.HasPrefix(word, "0x") {
			if value, err := strconv.ParseInt(word[2:], 16, 0); err == nil {
				return &Token{Op_Hex, word, byte(value)}, nil
			} else {
				return nil, err
			}
		}

		// -- Int
		if value, err := strconv.ParseInt(word, 10, 0); err == nil {
			return &Token{Op_Int, word, byte(value)}, nil
		} else {
			return nil, err
		}
	}
}
