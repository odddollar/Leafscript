package supporting

import (
	"strings"
)

func createVar(variableArray []string) {
	// check if exists
	if !checkIfVariableExists(variableArray[1]) {
		variables = append(variables, variable{
			Name: variableArray[1],
			Type: checkIfVariableIsString(variableArray[3]),
		})
	}

	// loop through variables to find right one
	for x := 0; x < len(variables); x++ {
		if variables[x].Name == variableArray[1] {
			// check if math is being performed or input is being requested
			if variableArray[3] != "math" && variableArray[3] != "input" && variableArray[3] != "string" && variableArray[3] != "concat" && variableArray[3] != "inputint" {
				// if math not being performed, assign value to variable
				// remove "" if string
				if variables[x].Type == "num" {
					variables[x].Value = strings.Join(variableArray[3:], "")
				} else {
					variables[x].Value = strings.ReplaceAll(strings.Join(variableArray[3:], " "), "\"", "")
				}
				break
			} else if variableArray[3] == "math" {
				// if math is being performed, run function and assign result to variable
				result := doMath(variableArray[4])
				variables[x].Value = result
				break
			} else if variableArray[3] == "input" {
				// check the string provided after "input" to print out
				result := input(strings.ReplaceAll(strings.Join(variableArray[4:], " "), "\"", ""))
				variables[x].Value = result
				break
			} else if variableArray[3] == "inputint" {
				// check the string provided after "input" to print out
				result := input(strings.ReplaceAll(strings.Join(variableArray[4:], " "), "\"", ""))
				variables[x].Value = result
				break
			} else if variableArray[3] == "string" {
				// check if another string is being assigned
				result := doString(variableArray[4])
				variables[x].Value = result
				break
			} else if variableArray[3] == "concat" {
				result := doConcat(strings.Join(variableArray[4:], " "))
				variables[x].Value = result
				break
			}
		}
	}
}

func checkIfVariableIsString(valueToCheck string) string {
	if strings.Contains(valueToCheck, "\"") || (valueToCheck == "input") || strings.Contains(valueToCheck, "string") || strings.Contains(valueToCheck, "concat") {
		return "str"
	} else {
		return "num"
	}
}

func checkIfVariableExists(variableToTest string) bool {
	found := false
	for x := 0; x < len(variables); x++ {
		if variables[x].Name == variableToTest {
			found = true
			break
		}
	}
	if found == true {
		return true
	} else {
		return false
	}
}
