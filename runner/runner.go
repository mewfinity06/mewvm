package runner

// import (
// "errors"
// "fmt"
//
// "github.com/mewfinity06/mewvm/lexer"
// )
//
// var (
// ErrorStackUnderflow = errors.New("stack underflow")
// ErrorStackOverflow  = errors.New("stack overflow")
// )
//
// const StackSize = 8
//
// type Register int
//
// const (
// Reg_A Register = iota
// Reg_B
// Reg_C
//
// Reg_Math
//
// RegCount
// )
//
// type Runner struct {
// Program   lexer.Program
// Stack     [StackSize]int
// Registers [RegCount]Register
//
// sp int
// }
//
// func NewRunner(program lexer.Program) Runner {
// return Runner{program, [StackSize]int{}, [RegCount]Register{}, 0}
// }
//
// func (r *Runner) Run() (int, error) {
// i := 0
// for ; i < len(r.Program); i++ {
// switch inst := r.Program[i]; inst {
// case int(lexer.TK_Add):
// if r.sp < 2 {
// return i, ErrorStackUnderflow
// }
//
// Get and reset a
// a := r.Stack[r.sp]
// r.Stack[r.sp] = 0
// r.sp -= 1
//
// Get and reset b
// b := r.Stack[r.sp]
// r.Stack[r.sp] = 0
// r.sp -= 1
//
// Set result
// r.Stack[r.sp] = b + a
//
// case int(lexer.TK_Push):
// if r.sp+1 >= StackSize { // r.sp+1 is for zero index
// return i, ErrorStackOverflow
// }
// i += 1
// op := r.Program[i]
// r.Stack[r.sp] = op
// r.sp += 1
//
// case int(lexer.TK_Pop):
// case int(lexer.TK_Eof): // eof does nothing
// default:
// return i, fmt.Errorf("unhandled inst: 0x%X", inst)
// }
// }
// return i, nil
// }
//
