package forthc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestParser(t *testing.T) {
	input := "5 10 +"
	var errBuf bytes.Buffer
	prog, err := Parser.ParseString("test.f", input, participle.Trace(&errBuf))
	if err != nil {
		fmt.Printf("\n%s", &errBuf)
		t.Fatal(err)
	}
	data, _ := json.MarshalIndent(prog, "", "  ")
	fmt.Printf("%s\n", string(data))
}

func TestGenerator(t *testing.T) {
	input := `: add3 + + ;
: add3_mul2 add3 2 * ;
1 2 3 add3_mul2`
	var errBuf bytes.Buffer
	prog, err := Parser.ParseString("test.f", input, participle.Trace(&errBuf))
	if err != nil {
		fmt.Printf("\n%s", errBuf.String())
		t.Fatal(err)
	}
	cd := NewGenerator()
	f, err := os.Create("test.asm")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	err = cd.GenerateFromProgram(prog, f)
	if err != nil {
		t.Fatal(err)
	}
}
