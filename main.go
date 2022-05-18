package main

import (
	"Leafscript/supporting"
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

func main() {
	// create argparse
	parser := argparse.NewParser("Leafscript", "The backend of the Leafscript programming language")
	run := parser.String("r", "run", &argparse.Options{Help: "The path to the .lfs file to run"})
	debug := parser.Flag("d", "debug", &argparse.Options{Default: false, Help: "Include to run program in debug mode"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
	}

	// ensure correct file format is used
	if strings.Contains(*run, ".lfs") {
		// parse file to language
		fileLines := supporting.ParseFile(*run)

		if *debug {
			supporting.Lex(fileLines, true)
		} else {
			supporting.Lex(fileLines, false)
		}
	} else {
		fmt.Println("Invalid file format. Please use the format .lfs")
	}
}
