package main

import (
	"github.com/odddollar/Leafscript/supporting"
	"flag"
	"fmt"
	"strings"
)

func main() {
	//fileLines := supporting.ParseFile("test.lfs")
	//supporting.Lex(fileLines)

	file := flag.String("run", "", "Directory of file to run")
	debug := flag.Bool("debug", false, "Run program with debug")
	flag.Parse()

	if strings.Contains(*file, ".lfs") {
		fileLines := supporting.ParseFile(*file)

		if *debug {
			supporting.Lex(fileLines, true)
		} else {
			supporting.Lex(fileLines, false)
		}
	} else {
		fmt.Println("Invalid file format. Please use the format .lfs")
	}
}
