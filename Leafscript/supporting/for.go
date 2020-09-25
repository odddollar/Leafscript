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
	} else {
		variables[getVariablePosition(iterVar)].Value = "0"
	}

	// get position of iteration variable in variable array
	pos := getVariablePosition(iterVar)

	var lastValue int

	// main for loop
	for x := 0; x < iterations; x++ {
		// check if the iteration variable has been reset, if so break continue loop from that point
		// not entirely sure why this works
		lastValue, _ = strconv.Atoi(variables[pos].Value)
		if lastValue+1 != x {
			x = lastValue
		} else {
			variables[pos].Value = strconv.Itoa(x)
		}
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