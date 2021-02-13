package supporting

func doString(variableName string) string {
	// create value to return
	var value string

	// iterate through variable list
	for x := 0; x < len(variables); x++ {
		// check if names match
		if variables[x].Name == variableName && variables[x].Type == "str" {
			value = variables[x].Value
			break
		}
	}

	return value
}
