package supporting

import (
	"fmt"
	"strings"
)

type variable struct {
	Name  string
	Type  string
	Value string
}

var variables = []variable{}
var globalDebug = false

func Lex(lines [][]string, debug bool) {
	//fmt.Println(lines)
	globalDebug = debug

	// iterate through lines on file
	for x := 0; x < len(lines); x++ {
		// run function based on keyword
		if strings.TrimSpace(lines[x][0]) == "var" && lines[x][2] == "=" {
			createVar(lines[x])
		} else if strings.TrimSpace(lines[x][0]) == "print" {
			customPrint(lines[x])
		} else if strings.TrimSpace(lines[x][0]) == "for" {
			// count number of tabs
			tabs := strings.Count(lines[x][0], "\t")

			forLines := [][]string{}
			// parse lines into for loop then skip lines in main lexer
			i := x + 1
			y := x
			for i < len(lines) {
				// break out of loop
				if strings.TrimSpace(lines[i][0]) == "endfor" && strings.Count(lines[i][0], "\t") == tabs {
					x = i
					break
				}

				// append current line to list of lines to parse to for loop
				forLines = append(forLines, lines[i])

				i += 1
			}

			// parse lines to for loop lexer
			forLoop(forLines, lines[y][1], strings.Join(lines[y][2:], " "))
		} else if strings.TrimSpace(lines[x][0]) == "if" {
			// count number of tabs
			tabs := strings.Count(lines[x][0], "\t")

			ifLines := [][]string{}
			elseLines := [][]string{}
			isElse := false
			// parse lines into if statement then skip lines in main lexer
			i := x + 1
			y := x
			for i < len(lines) {
				// break out of loop
				if (strings.TrimSpace(lines[i][0]) == "endif" || strings.TrimSpace(lines[i][0]) == "else") && strings.Count(lines[i][0], "\t") == tabs {
					// check if statement is present
					if strings.TrimSpace(lines[i][0]) == "else" && strings.Count(lines[i][0], "\t") == tabs {
						isElse = true
						elseLines = getElseLines(i, lines, tabs)
						x = i
					}
					x = i
					break
				}

				// append current line to list of lines to parse to if statement
				ifLines = append(ifLines, lines[i])

				i += 1
			}

			// parse lines to if statement lexer based on whether else is present
			if isElse == false {
				ifStatement(ifLines, [][]string{}, strings.Join(lines[y][1:], " "))
			} else {
				ifStatement(ifLines, elseLines, strings.Join(lines[y][1:], " "))
				for s := x; s < len(lines); s++ {
					if strings.TrimSpace(lines[s][0]) == "endif" && strings.Count(lines[s][0], "\t") == tabs {
						x = s
						break
					}
				}
			}
		}

		// print variable list if debug mode is true
		if debug {
			fmt.Println(variables)
		}
	}
}

func getElseLines(currentLine int, lines [][]string, tabs int) [][]string {
	currentLine += 1
	returnLines := [][]string{}

	// check if reached the end of if statement
	for currentLine < len(lines) {
		if strings.TrimSpace(lines[currentLine][0]) == "endif" && strings.Count(lines[currentLine][0], "\t") == tabs {
			break
		}

		returnLines = append(returnLines, lines[currentLine])

		currentLine += 1
	}

	return returnLines
}
