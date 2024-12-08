package main

import (
	"fmt"
	"strconv"
	"strings"

	adventofcode "github.com/broothie/advent-of-code"
)

type Equation struct {
	value    int
	operands []int
}

func (e Equation) canBeMadeTrue(operators string) bool {
	for _, ops := range generateOperatorSets(operators, len(e.operands)-1) {
		value := e.operands[0]
		for index, op := range ops {
			rightOperand := e.operands[index+1]

			if op == '+' {
				value += rightOperand
			} else if op == '*' {
				value *= rightOperand
			} else if op == '|' {
				parsedValue, err := strconv.Atoi(strconv.Itoa(value) + strconv.Itoa(rightOperand))
				if err != nil {
					panic(fmt.Sprintf("unable to parse concatenated value: %v", err))
				}

				value = parsedValue
			}
		}

		if value == e.value {
			return true
		}
	}

	return false
}

func generateOperatorSets(operators string, operatorCount int) []string {
	if operatorCount == 0 {
		return []string{""}
	}

	var operatorSets []string
	for _, operatorSet := range generateOperatorSets(operators, operatorCount-1) {
		for _, operator := range operators {
			newOperatorSet := operatorSet + string(operator)
			operatorSets = append(operatorSets, newOperatorSet)
		}
	}

	return operatorSets
}

func parse(input string) ([]Equation, error) {
	var equations []Equation
	for _, equationLine := range strings.Split(input, "\n") {
		equationSplit := strings.SplitN(equationLine, ": ", 2)
		valueString, operandsString := equationSplit[0], equationSplit[1]

		value, err := strconv.Atoi(valueString)
		if err != nil {
			return nil, err
		}

		var operands []int
		for _, operandString := range strings.Fields(operandsString) {
			operand, err := strconv.Atoi(operandString)
			if err != nil {
				return nil, err
			}

			operands = append(operands, operand)
		}

		equations = append(equations, Equation{
			value:    value,
			operands: operands,
		})
	}

	return equations, nil
}

func main() {
	adventofcode.Part1(2024, 7, 3749, func(input string) (int, error) {
		equations, err := parse(input)
		if err != nil {
			return 0, err
		}

		sum := 0
		for _, equation := range equations {
			if equation.canBeMadeTrue("+*") {
				sum += equation.value
			}
		}

		return sum, nil
	})

	adventofcode.Part2(2024, 7, 11387, func(input string) (int, error) {
		equations, err := parse(input)
		if err != nil {
			return 0, err
		}

		sum := 0
		for _, equation := range equations {
			if equation.canBeMadeTrue("+*|") {
				sum += equation.value
			}
		}

		return sum, nil
	})
}
