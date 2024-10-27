package forthc

import (
	"fmt"
	"io"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var Lexer = lexer.MustStateful(lexer.Rules{
	"Root": []lexer.Rule{
		{Name: "whitespace", Pattern: `[ \r\t\n]+`},

		{Name: "Symbol", Pattern: `[a-zA-Z_]+`},
		{Name: "Integer", Pattern: `(-)?[0-9]+`},

		{Name: "+", Pattern: `\+`},
		{Name: "-", Pattern: `-`},
		{Name: "*", Pattern: `\*`},
		{Name: "/", Pattern: `\/`},
	},
})

var Parser = participle.MustBuild[Program](
	participle.Lexer(Lexer),
	participle.UseLookahead(1),
	participle.Union[Expression](
		IntegerNode{},
		BinOpNode{},
	),
)

/*
=== AST
*/
type Program struct {
	Expressions []Expression `parser:"@@+"`
}

type Expression interface{}

type IntegerNode struct {
	Value int32 `parser:"@Integer"`
}

type BinOpNode struct {
	Operation string `parser:"@('+'|'-'|'*'|'/')"`
}

/*
=== Code generator
*/
const Preamble = `init:
li sp, 0x10010000
j main

main:
`

type Codegen struct{}

func (cd *Codegen) GeneratePreamble(out io.Writer) {
	fmt.Fprint(out, Preamble)
}

func (cd *Codegen) GenerateFromProgram(prog *Program, out io.Writer) error {
	cd.GeneratePreamble(out)
	for _, node := range prog.Expressions {
		switch node := node.(type) {
		case IntegerNode:
			if node.Value < 0 {
				fmt.Fprintf(out, "li t0, -0x%x\n", node.Value*-1)
			} else {
				fmt.Fprintf(out, "li t0, 0x%x\n", node.Value)
			}
			fmt.Fprintf(out, "sw t0, 0(sp)\n")
			fmt.Fprintf(out, "addi sp, sp, 0x4\n")
		case BinOpNode:
			fmt.Fprintf(out, "li t0, 0x4\n")
			fmt.Fprintf(out, "sub sp, sp, t0\n")
			fmt.Fprintf(out, "lw t2, 0(sp)\n")
			fmt.Fprintf(out, "sub sp, sp, t0\n")
			fmt.Fprintf(out, "lw t1, 0(sp)\n")
			switch node.Operation {
			case "+":
				fmt.Fprintf(out, "add t0, t1, t2\n")
			case "-":
				fmt.Fprintf(out, "sub t0, t1, t2\n")
			case "*":
				fmt.Fprintf(out, "mul t0, t1, t2\n")
			case "/":
				fmt.Fprintf(out, "div t0, t1, t2\n")
			default:
				return fmt.Errorf("unexpected binary operation: %s", node.Operation)
			}
			fmt.Fprintf(out, "sw t0, 0(sp)\n")
			fmt.Fprintf(out, "addi sp, sp, 0x4\n")
		default:
			return fmt.Errorf("receive unexpected node: %T", node)
		}
	}
	return nil
}
