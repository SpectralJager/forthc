package forthc

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestGenerator(t *testing.T) {
	input := `variable counter
: count_to_10
	begin
		counter @ 1 + counter !
		counter @ 10 =		
	until
;
0 counter !
count_to_10`
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
