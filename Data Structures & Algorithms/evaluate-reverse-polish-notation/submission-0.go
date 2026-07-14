func evalRPN(tokens []string) int {
	stack := []int{}

	for _, ch := range tokens {
		if !isOperator(ch) {
			operand, err := strconv.Atoi(ch)

			if err != nil {
				panic("Invalid token")
			}

			stack = append(stack, operand)
			continue
		}

		opd2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		opd1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, calc(ch, opd1, opd2))
	}

	return stack[0]
}

func isOperator(val string) bool {
	switch val {
	case "+", "-", "*", "/":
		return true
	}
	return false
}

func calc(opr string, opd1, opd2 int) int {
	switch opr {
	case "+":
		return opd1 + opd2
	case "-":
		return opd1 - opd2
	case "*":
		return opd1 * opd2
	case "/":
		if opd2 == 0 {
			panic("division by zero")
		}
		return opd1 / opd2
	}
	return 0
}