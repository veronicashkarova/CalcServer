package CalcServer

import (
	"errors"
	"strconv"
	"strings"
)

var Priority = map[string]int{
	")": 0,
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func toPostFix(expression string) ([]string, error) {
	symbols := strings.Split(expression, "")
	postfixExpr := []string{}

	stack := []string{}

    isLastNumber := false

	for _, symbol := range symbols {
		if _, err := strconv.ParseFloat(symbol, 64); err == nil {
			postfixExpr = append(postfixExpr, symbol)
            isLastNumber = true
		} else {

			switch {
			case symbol[0] == '(':
				stack = append(stack, symbol)
			case symbol[0] == ')':
				for len(stack) > 0 && stack[len(stack)-1] != "(" {
					postfixExpr = append(postfixExpr, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
                stack = stack[:len(stack)-1]
            
			default:
                if !isLastNumber {
                    return []string{}, errors.New("неправильное выражение")
                }
				if _, ok := Priority[symbol]; ok {
					for len(stack) > 0 && Priority[stack[0]] >= Priority[symbol] {
						postfixExpr = append(postfixExpr, stack[len(stack)-1])
						stack = stack[:len(stack)-1]

					}
					stack = append(stack, symbol)
                    isLastNumber = false
				} else {
					return []string{}, errors.New("неправильное выражение")
				}
			}
		}
	}

	for i := range stack {
		postfixExpr = append(postfixExpr, stack[len(stack)-1-i])
	}
    if len(postfixExpr)%2 == 0 {
        return []string{}, errors.New("неправильное выражение")
    }
	return postfixExpr, nil
}

func CalcPostfix(postfixExpr []string) (float64, error) {
	locals := []float64{}
	var counter = 0

	for _, symbol := range postfixExpr {
		if tokenNum, err := strconv.ParseFloat(symbol, 64); err == nil {
			locals = append(locals, tokenNum)
		} else if _, ok := Priority[symbol]; ok {
			counter += 1
			var second float64 = 0
			var first float64 = 0
			if len(locals) > 0 {
				second = locals[len(locals)-1]
				locals = locals[:len(locals)-1]
			}
			if len(locals) > 0 {
				first = locals[len(locals)-1]
				locals = locals[:len(locals)-1]
			}
			switch symbol {
			case "+":
				locals = append(locals, first+second)
			case "-":
				locals = append(locals, first-second)
			case "*":
				locals = append(locals, first*second)
			case "/":
				if second == 0 {
					return 0, errors.New("деление на ноль")
				} else {
					locals = append(locals, first/second)
				}
			}
		}
	}
	return locals[len(locals)-1], nil
}

func Calc(expression string) (float64, error) {

    if len(expression) == 0 {
        return 0, errors.New("пустая строка")
    }
	postfix, err := toPostFix(expression)

	if err != nil {
		return 0, err
	} else {
		result, err2 := CalcPostfix(postfix)
		if err2 != nil {
			return 0, err2
		}
		return result, nil
	}
}