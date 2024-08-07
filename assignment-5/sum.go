package main

import (
	"errors"
	"fmt"
	"log"
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

func Sum(numbers []int, maxDigits int) int {
	validateLengthErrors := validateLengthOfArray(numbers)

	if validateLengthErrors != nil {
		handleError(validateLengthErrors)
	}

	digitValidationErrors := validateNumberDigits(numbers, maxDigits)

	if digitValidationErrors != nil {
		handleError(digitValidationErrors)
	}

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var singleSum, doubleSum, tripleSum int

	singleDigits := [3]int{7, 8, 9}
	doubleDigits := [3]int{55, 22, 17}
	tripleDigits := [3]int{333, 555, 888}

	singleSum = Sum(singleDigits[:], 1)
	doubleSum = Sum(doubleDigits[:], 2)
	tripleSum = Sum(tripleDigits[:], 3)

	fmt.Fprintln(os.Stdout, "Single Digits Sum: ", singleSum)
	fmt.Fprintln(os.Stdout, "Double Digits Sum: ", doubleSum)
	fmt.Fprintln(os.Stdout, "Triple Digits Sum: ", tripleSum)

	combinedSums := []int{singleSum, doubleSum, tripleSum}
	sumOfSums := Sum(combinedSums, -1)

	fmt.Fprintln(os.Stdout, "Sum of Sums: ", sumOfSums)
}
