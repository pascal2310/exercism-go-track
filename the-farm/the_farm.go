package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(calc FodderCalculator, amountOfCows int) (float64, error) {
	total, err := calc.FodderAmount(amountOfCows)
	if err != nil {
		return 0, err
	}
	perCow := total / float64(amountOfCows)

	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return perCow * factor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, amountOfCows int) (float64, error) {
	if amountOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calc, amountOfCows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	numberOfCows  int
	customMessage string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.numberOfCows, err.customMessage)
}

func ValidateNumberOfCows(amountOfCows int) error {
	if amountOfCows < 0 {
		return &InvalidCowsError{
			amountOfCows,
			"there are no negative cows",
		}
	}
	if amountOfCows == 0 {
		return &InvalidCowsError{
			amountOfCows,
			"no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
