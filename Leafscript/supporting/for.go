package supporting

import (
	"strconv"
	"strings"
)

func forLoop(lines [][]string, iterVar string, iter string) {
	// check if iter is using variable or just number
	y := strings.Split(iter, " ")
	var iterations int
	if y[0] == "math" {
		iterations, _ = strconv.Atoi(doMath(y[1]))
	} else {
		iterations, _ = strconv.Atoi(y[0])
	}

	// create iteration variable
	if !checkIfVariableExists(iterVar) {
		variables = append(variables, variable{
			Name:  iterVar,
			Type:  "num",
		})
	}

	// get position of iteration variable in variable array
	pos := getVariablePosition(iterVar)

	// main for loop
	for x := 0; x < iterations; x++ {
		variables[pos].Value = strconv.Itoa(x)
		Lex(lines)
	}
}

func getVariablePosition(variableToCheck string) int {
	var pos int
	for x := 0; x < len(variables); x++ {
		if variables[x].Name == variableToCheck {
			pos = x
			break
		}
	}

	return pos
}