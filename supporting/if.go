package supporting

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strings"
)

func ifStatement(lines [][]string, falseLines [][]string, conditional string) {
	var conditionalSplitLeft, conditionalSplitRight []string
	var operator, final string

	// split conditional statement into left and right to parse to function
	conditionalSplit := strings.Split(conditional, " ")
	for x := 0; x < len(conditionalSplit); x++ {
		if conditionalSplit[x] == "==" || conditionalSplit[x] == "!=" || conditionalSplit[x] == "<=" || conditionalSplit[x] == ">=" || conditionalSplit[x] == "<" || conditionalSplit[x] == ">" {
			conditionalSplitLeft = conditionalSplit[:x]
			conditionalSplitRight = conditionalSplit[x+1:]
			operator = conditionalSplit[x]
			break
		}
	}

	// parse left and right parts of conditional to replacing function
	final += replaceConditionalPart(conditionalSplitLeft)
	final += operator
	final += replaceConditionalPart(conditionalSplitRight)

	// parse in expression for evaluation
	exp, err := govaluate.NewEvaluableExpression(final)
	if err != nil {
		panic(err)
	}
	result, _ := exp.Evaluate(nil)

	// main lexer to be run if expression is true
	if result == true {
		// parse lines into main lexer
		Lex(lines, globalDebug)
	} else {
		Lex(falseLines, globalDebug)
	}
}

func replaceConditionalPart(half []string) string {
	// check what to return
	if half[0] == "math" {
		return doMath(half[1])
	} else if half[0] == "string" {
		return fmt.Sprintf("\"%v\"", doString(half[1]))
	} else if strings.Contains(half[0], "\"") {
		return fmt.Sprintf("\"%v\"", strings.ReplaceAll(strings.Join(half[0:], " "), "\"", ""))
	} else {
		return half[0]
	}
}
