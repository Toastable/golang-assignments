package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bearbin/go-age"
)

func CalculateAge(dateOfBirth string) (int, error) {
	convertedDateOfBirth, err := convertToDate(dateOfBirth)

	if err != nil {
		return 0, err
	}

	ageOfPerson := age.Age(convertedDateOfBirth)

	return ageOfPerson, nil
}

func convertToDate(dateOfBirth string) (time.Time, error) {
	layout := "01/02/2006"

	output, err := time.Parse(layout, dateOfBirth)

	if err != nil {
		return time.Time{}, err
	}

	return output, nil
}

func main() {
	fmt.Println("Enter Dates of Birth (DD/MM/YYYY)")
	fmt.Println("|----------------------|")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		err := scanner.Err()

		if err != nil {
			log.Println("Invalid input")
		}

		message, genErr := CalculateAge(input)

		if genErr != nil {
			log.Println("Input could not be converted into a valid date")
		} else {
			fmt.Fprintln(os.Stdout, "The person's age is:", message)
		}
	}
}
