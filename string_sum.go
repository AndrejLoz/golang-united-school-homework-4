package string_sum

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")

	errorNotBadOperation = errors.New("bad operation in expecting")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
// For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var tmpSlice []string
	var startFor int = 0
	var valArr [2]string

	input = strings.TrimSpace(input)
	if input == "" {
		return "", errorEmptyInput
	}

	a := regexp.MustCompile("[\\+-]")
	numOperand := len(a.Split(input, -1))
	if (string(input[0]) == "-") || (string(input[0]) == "+") {
		numOperand--
		startFor = 1
	}
	tmpSlice = a.Split(input, -1)
	if numOperand == 1 || numOperand > 2 {
		return "", errorNotTwoOperands
	}

	for i := startFor; i < len(tmpSlice); i++ {
		_, err := strconv.Atoi(strings.TrimSpace(tmpSlice[i]))
		if err != nil {
			return "", err
		}
	}

	j := 0
	for i := 0; i < len(input); i++ {
		if (string(input[i]) == "+" || string(input[i]) == "-") && (i > 0) {
			j = 1
		}
		if string(input[i]) != " " {
			valArr[j] += string(input[i])
		}
	}

	val_1_int, _ := strconv.Atoi(valArr[0])
	val_2_int, _ := strconv.Atoi(valArr[1])
	sumVal := val_1_int + val_2_int

	return strconv.Itoa(sumVal), nil
}
