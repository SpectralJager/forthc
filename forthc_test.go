package forthc

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestGenerator(t *testing.T) {
	input := `\3 4 < 20 30 < and
\3 4 < 20 30 or
\3 4 < invert
-1 -1 and`
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
