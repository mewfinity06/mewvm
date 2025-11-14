package runner

import (
	"github.com/mewfinity06/mewvm/lexer"
)

type Runner struct {
	lexer.Program
	cur int
}

func NewRunner(program lexer.Program) Runner {
	return Runner{program, 0}
}

func (r Runner) Run(debug bool) (int, error) {
	return 0x0, nil
}
