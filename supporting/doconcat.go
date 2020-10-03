package supporting

import (
	"strings"
)

func doConcat(toConcat string) string {
	// split inputted string into array to operate on and create final string
	split := strings.Split(toConcat, "&")
	var final string

	// iterate through to remove whitespace
	for i := 0; i < len(split); i++ {
		split[i] = strings.TrimSpace(split[i])
	}

	// iterate through array and perform actions
	for i := 0; i < len(split); i++ {
		if strings.Contains(split[i], "\"") {
			final += strings.ReplaceAll(split[i], "\"", "")
		} else if strings.Contains(split[i], "string") {
			final += doString(strings.Split(split[i], " ")[1])
		} else if strings.Contains(split[i], "math") {
			final += doMath(strings.Split(split[i], " ")[1])
		}
	}

	return final
}
