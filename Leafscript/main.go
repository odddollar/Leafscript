package main

import (
	"./supporting"
)

func main() {
	fileLines := supporting.ParseFile("test.lfs")
	supporting.Lex(fileLines)

	/*file := flag.String("file_path", "", "Directory of file to run")
	flag.Parse()

	if strings.Contains(*file, ".lfs") {
		fileLines := supporting.ParseFile(*file)

		supporting.Lex(fileLines)
	} else {
		fmt.Println("Invalid file format. Please use the format .lfs")
	}*/
}
