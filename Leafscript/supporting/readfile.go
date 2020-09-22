package supporting

import (
	"bufio"
	"os"
	"strings"
)

func ParseFile(directory string) [][]string {
	// open project file
	file, err := os.Open(directory)
	if err != nil {
		panic(err)
	}

	// create array of file lines
	lines := [][]string{}

	// create scanner to read file
	scanner := bufio.NewScanner(file)

	// read through file lines
	for scanner.Scan() {
		line := scanner.Text()
		// check if line commented out
		if !strings.Contains(line, "//") {
			line = strings.TrimSpace(line)
			lineSplit := strings.Split(line, " ")
			lines = append(lines, lineSplit)
		}
	}

	return lines
}
