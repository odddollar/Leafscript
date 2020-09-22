package supporting

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strconv"
)

func doMath(expression string) string {
	// create dictionary to use as parameters
	params := make(map[string]interface{}, 8)
	for x := 0; x < len(variables); x++ {
		if variables[x].Type == "num" {
			params[variables[x].Name], _ = strconv.ParseFloat(variables[x].Value, 32)
		}
	}

	// create new expression and evaluate
	exp, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		panic(err)
	}
	result, err := exp.Evaluate(params)

	return fmt.Sprintf("%v", result)
}