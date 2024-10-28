package forthc

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/google/uuid"
)

var Lexer = lexer.MustStateful(lexer.Rules{
	"Root": []lexer.Rule{
		{Name: "whitespace", Pattern: `[ \r\t\n]+`},

		{Name: "If", Pattern: `if`},
		{Name: "Then", Pattern: `then`},
		{Name: "Else", Pattern: `else`},
		{Name: "And", Pattern: `and`},
		{Name: "Or", Pattern: `or`},
		{Name: "Invert", Pattern: `invert`},
		{Name: "Symbol", Pattern: `[a-zA-Z_]+[a-zA-Z_0-9]*[\?<>(<>)=(<=)(>=)]?`},
		{Name: "Integer", Pattern: `(-)?[0-9]+`},

		{Name: "+", Pattern: `\+`},
		{Name: "-", Pattern: `-`},
		{Name: "*", Pattern: `\*`},
		{Name: "/", Pattern: `\/`},
		{Name: "<>", Pattern: `<>`},
		{Name: "<", Pattern: `<`},
		{Name: ">", Pattern: `>`},
		{Name: "<=", Pattern: `<=`},
		{Name: ">=", Pattern: `>=`},
		{Name: "=", Pattern: `=`},
		{Name: ":", Pattern: `:`},
		{Name: ";", Pattern: `;`},
	},
})

var Parser = participle.MustBuild[Program](
	participle.Lexer(Lexer),
	participle.UseLookahead(1),
	participle.Union[Expression](
		IntegerNode{},
		SymbolNode{},
		BinOpNode{},
		UnOpNode{},
		SymbolDefNode{},
	),
	participle.Union[DefinitionExpression](
		IntegerNode{},
		SymbolNode{},
		BinOpNode{},
		UnOpNode{},
		IfThenElseNode{},
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

type SymbolNode struct {
	Value string `parser:"@Symbol"`
}

type BinOpNode struct {
	Operation string `parser:"@('+'|'-'|'*'|'/'|'<'|'>'|'<='|'>='|'='|'<>'|'and'|'or')"`
}

type UnOpNode struct {
	Operations string `parser:"@('invert')"`
}

type DefinitionExpression interface{}

type SymbolDefNode struct {
	Symbol string                 `parser:"':' @Symbol"`
	Body   []DefinitionExpression `parser:"@@+ ';'"`
}

type IfThenElseNode struct {
	ThenBody []DefinitionExpression `parser:"'if' @@+ "`
	ElseBody []DefinitionExpression `parser:"('else' @@+)? 'then'"`
}

/*
=== Code generator
*/
const Preamble = `j .init
.init:
li sp, 0x10010000
j .main

.main:
`

type Codegen struct {
	Environment map[string]string
}

func NewGenerator() *Codegen {
	return &Codegen{
		Environment: map[string]string{
			"dup": `addi sp, sp, -0x4
lw t0, 0(sp)
addi sp, sp, 0x4
sw t0, 0(sp)
addi sp, sp, 0x4
`,
			"swap": `addi sp, sp, -0x4
lw t0, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
sw t0, 0(sp)
addi sp, sp, 0x4
sw t1, 0(sp)
addi sp, sp, 0x4
`,
			"drop": `addi sp, sp, -0x4`,
		},
	}
}

func (cd *Codegen) GeneratePreamble(out io.Writer) {
	fmt.Fprint(out, Preamble)
}

func (cd *Codegen) GenerateFromProgram(prog *Program, out io.Writer) error {
	cd.GeneratePreamble(out)
	for _, node := range prog.Expressions {
		err := cd.Generate(node, out)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cd *Codegen) Generate(node any, out io.Writer) error {
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
		lbl := strings.ReplaceAll(uuid.NewString(), "-", "_")
		fmt.Fprintf(out, "addi sp, sp, -0x4\n")
		fmt.Fprintf(out, "lw t2, 0(sp)\n")
		fmt.Fprintf(out, "addi sp, sp, -0x4\n")
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
		case "<":
			fmt.Fprintf(out, "slt t0, t1, t2\n")
			fmt.Fprintf(out, "neg t0, t0\n")
		case ">":
			fmt.Fprintf(out, "slt t0, t2, t1\n")
			fmt.Fprintf(out, "neg t0, t0\n")
		case "<=":
			fmt.Fprintf(out, "li t0, 1\n")
			fmt.Fprintf(out, "beq t1, t2, .%s\n", lbl)
			fmt.Fprintf(out, "slt t0, t1, t2\n")
			fmt.Fprintf(out, ".%s:\n", lbl)
			fmt.Fprintf(out, "neg t0, t0\n")
		case ">=":
			fmt.Fprintf(out, "li t0, 1\n")
			fmt.Fprintf(out, "beq t1, t2, .%s\n", lbl)
			fmt.Fprintf(out, "slt t0, t2, t1\n")
			fmt.Fprintf(out, ".%s:\n", lbl)
			fmt.Fprintf(out, "neg t0, t0\n")
		case "=":
			fmt.Fprintf(out, "li t0, 1\n")
			fmt.Fprintf(out, "beq t1, t2, .%s\n", lbl)
			fmt.Fprintf(out, "li t0, 0\n")
			fmt.Fprintf(out, ".%s:\n", lbl)
			fmt.Fprintf(out, "neg t0, t0\n")
		case "<>":
			fmt.Fprintf(out, "li t0, 1\n")
			fmt.Fprintf(out, "bne t1, t2, .%s\n", lbl)
			fmt.Fprintf(out, "li t0, 0\n")
			fmt.Fprintf(out, ".%s:\n", lbl)
			fmt.Fprintf(out, "neg t0, t0\n")
		default:
			return fmt.Errorf("unexpected binary operation: %s", node.Operation)
		}
		fmt.Fprintf(out, "sw t0, 0(sp)\n")
		fmt.Fprintf(out, "addi sp, sp, 0x4\n")
	case SymbolNode:
		body, ok := cd.Environment[node.Value]
		if !ok {
			return fmt.Errorf("undefined symbol: %s", node.Value)
		}
		fmt.Fprint(out, body)
	case SymbolDefNode:
		var body bytes.Buffer
		for _, exp := range node.Body {
			err := cd.Generate(exp, &body)
			if err != nil {
				return err
			}
		}
		cd.Environment[node.Symbol] = body.String()
	case IfThenElseNode:
		var thenBody, elseBody bytes.Buffer
		for _, exp := range node.ThenBody {
			err := cd.Generate(exp, &thenBody)
			if err != nil {
				return err
			}
		}
		for _, exp := range node.ElseBody {
			err := cd.Generate(exp, &elseBody)
			if err != nil {
				return err
			}
		}
		lbl := strings.ReplaceAll(uuid.NewString(), "-", "_")
		fmt.Fprintf(out, "addi sp, sp, -0x4\n")
		fmt.Fprintf(out, "lw t0, 0(sp)\n")
		fmt.Fprintf(out, "beq t0, zero, .%s_else\n", lbl)
		fmt.Fprint(out, thenBody.String())
		fmt.Fprintf(out, "j .%s\n", lbl)
		fmt.Fprintf(out, ".%s_else:\n", lbl)
		fmt.Fprint(out, elseBody.String())
		fmt.Fprintf(out, ".%s:\n", lbl)
	default:
		return fmt.Errorf("receive unexpected node: %T", node)
	}
	return nil
}
