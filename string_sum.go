package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func parseNumber(input string, lastDelimiter byte) (output int, err error) {
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, err
	}
	if lastDelimiter == '-' {
		num = -num
	}
	return num, nil
}

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	var result int = 0

	trimmedInput := strings.TrimSpace(input)
	// try find equation delimeter
	countDelimeters := 0
	idxFound := 0
	var lastDelimiter byte = '+'
	for i, il := 0, len(trimmedInput); i < il; i++ {
		if trimmedInput[i] == '+' || trimmedInput[i] == '-' {
			if i != 0 {
				tmp := trimmedInput[idxFound:i]
				num, err := parseNumber(strings.TrimSpace(tmp), lastDelimiter)
				if err != nil {
					return "", fmt.Errorf("Invalid number: %w", err)
				}
				result += num
				countDelimeters++
			}

			lastDelimiter = trimmedInput[i]
			idxFound = i + 1
		}
	}

	if idxFound != len(trimmedInput) {
		tmp := trimmedInput[idxFound:]
		num, err := parseNumber(strings.TrimSpace(tmp), lastDelimiter)
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}
		result += num
		countDelimeters++
	}

	if countDelimeters != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	return strconv.Itoa(result), nil
}
