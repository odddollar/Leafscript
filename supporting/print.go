package supporting

import (
	"fmt"
	"strings"
)

func customPrint(variableArray []string) {
	// check if printing string
	if strings.Contains(variableArray[1], "\"") {
		fmt.Println(strings.ReplaceAll(strings.Join(variableArray[1:], " "), "\"", ""))
	// check if doing math or printing string variable
	} else if variableArray[1] == "math" {
		fmt.Println(doMath(variableArray[2]))
	} else if variableArray[1] == "string" {
		fmt.Println(doString(variableArray[2]))
	} else if variableArray[1] == "concat" {
		fmt.Println(doConcat(strings.Join(variableArray[2:], " ")))
	}
}
