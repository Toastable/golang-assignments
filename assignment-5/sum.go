package main

import (
	"errors"
	"fmt"
	"os"
)

var errInvalidDigitsError = errors.New("invalid number of digits")
var errNoNegativeDigitsError = errors.New("negative numbers are not allowed")
var errTooManyNumbers = errors.New("no more than 3 numbers allowed")

func validateNumberDigits(numbers []int, numberOfDigits int) error {
	if numberOfDigits == -1 {
		return nil
	}

	for _, number := range numbers {
		if number < 0 {
			return errNoNegativeDigitsError
		}
		if numberOfDigits == 1 && number > 9 {
			return errInvalidDigitsError
		} else if numberOfDigits == 2 && number > 99 {
			return errInvalidDigitsError
		} else if numberOfDigits == 3 && number > 999 {
			return errInvalidDigitsError
		}
	}

	return nil
}

func validateLengthOfArray(numbers []int) error {
	if len(numbers) > 3 {
		return errTooManyNumbers
	}

	return nil
}

func Sum(numbers []int, maxDigits int) (int, error) {
	validateLengthErrors := validateLengthOfArray(numbers)

	if validateLengthErrors != nil {
		return 0, validateLengthErrors
	}

	digitValidationErrors := validateNumberDigits(numbers, maxDigits)

	if digitValidationErrors != nil {
		return 0, digitValidationErrors
	}

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum, nil
}

func main() {
	var singleSum, doubleSum, tripleSum int

	singleDigits := [3]int{7, 8, 9}
	doubleDigits := [3]int{55, 22, 17}
	tripleDigits := [3]int{333, 555, 888}

	singleSum, _ = Sum(singleDigits[:], 1)
	doubleSum, _ = Sum(doubleDigits[:], 2)
	tripleSum, _ = Sum(tripleDigits[:], 3)

	fmt.Fprintln(os.Stdout, "Single Digits Sum: ", singleSum)
	fmt.Fprintln(os.Stdout, "Double Digits Sum: ", doubleSum)
	fmt.Fprintln(os.Stdout, "Triple Digits Sum: ", tripleSum)

	combinedSums := []int{singleSum, doubleSum, tripleSum}
	sumOfSums, _ := Sum(combinedSums, -1)

	fmt.Fprintln(os.Stdout, "Sum of Sums: ", sumOfSums)
}
