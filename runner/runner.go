package runner

import (
	"fmt"

	"github.com/mewfinity06/mewvm/lexer"
	"github.com/mewfinity06/mewvm/packer"
)

type Runner struct {
	Prog packer.Program
}

func RunnerNew(prog packer.Program) Runner {
	return Runner{prog}
}

func (r Runner) Run() (int, error) {
	i := 0
	return i, nil
}

func (r Runner) PrintProgram() {
	for i := 0; i < len(r.Prog); i++ {
		switch inst := lexer.Instruction(r.Prog[i]); inst {
		case lexer.Inst_Push:
			fmt.Print("push ")
			i += 1
			op := r.Prog[i]
			fmt.Printf("0x%X\n", op)
		case lexer.Inst_Add:
			fmt.Println("add")
		case lexer.Inst_Pop:
			fmt.Print("pop ")
			i += 1
			switch reg := lexer.Register(r.Prog[i]); reg {
			case lexer.Reg_A:
				fmt.Println("A")
			default:
				fmt.Printf("unhandled register: 0x%X\n", reg.ToHex())
			}
		case lexer.Inst_EOF:
			fmt.Println("eof")
		default:
			fmt.Printf("unhandled instruction: 0x%d\n", inst)
		}
	}
}
