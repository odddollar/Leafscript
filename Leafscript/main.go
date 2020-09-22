package main

import (
	"./supporting"
)

func main() {
	//file := flag.String("file_path", "", "Directory of file to run")
	//flag.Parse()

	fileLines := supporting.ParseFile("test.lfs")
	//fileLines := supporting.ParseFile(*file)

	supporting.Lex(fileLines)
}
