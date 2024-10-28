package main

import (
	"flag"
	"log"
	"os"

	"github.com/SpectralJager/forthc"
	"github.com/alecthomas/participle/v2"
)

var (
	sourcePath = flag.String("src", "main.fth", "path to the source file")
	outPath    = flag.String("out", "main.asm", "path to the output file")
	help       = flag.Bool("help", false, "print help message")
)

func init() {
	flag.Parse()
}

func main() {
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	data, err := os.ReadFile(*sourcePath)
	if err != nil {
		log.Fatalf("couldn't read source file: %s", err.Error())
	}
	logFile, err := os.Create("logs.log")
	if err != nil {
		log.Fatalf("couldn't create logs file: %s", err.Error())
	}
	prog, err := forthc.Parser.ParseBytes(*sourcePath, data, participle.Trace(logFile))
	if err != nil {
		log.Fatalf("can't parse source file: %s", err.Error())
	}
	gen := forthc.NewGenerator()
	outFile, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("can't create output file: %s", err.Error())
	}
	err = gen.GenerateFromProgram(prog, outFile)
	if err != nil {
		log.Fatalf("can't generate output file: %s", err.Error())
	}
}
