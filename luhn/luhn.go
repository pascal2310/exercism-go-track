package luhn

import (
	"errors"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	withoutSpaces := strings.ReplaceAll(id, " ", "")
	if len(withoutSpaces) <= 1 {
		return false
	}

	intSlice, err := toIntSlice(withoutSpaces)
	if err != nil {
		return false
	}

	intSlice = doubleEverySecondDigitStartingFromRight(intSlice)

	return sum(intSlice)%10 == 0
}

func sum(input []int) int {
	var result int

	for i := range input {
		result += input[i]
	}

	return result
}

func toIntSlice(id string) ([]int, error) {
	var intSlice = make([]int, len(id))

	for i := range id {
		integer, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return nil, errors.New("not a valid integer")
		}

		intSlice[i] = integer
	}

	return intSlice, nil
}

func doubleEverySecondDigitStartingFromRight(input []int) []int {
	// i is len of input -2 so the first 'second' digit from the right, after that take steps of 2
	for i := len(input) - 2; i >= 0; i = i - 2 {
		newValue := input[i] * 2

		if newValue > 9 {
			newValue -= 9
		}
		input[i] = newValue
	}

	return input
}
