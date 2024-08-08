package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter Numbers")
	fmt.Println("|----------------------|")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		err := scanner.Err()

		if err != nil {
			log.Println("Invalid input")
		}

		message, genErr := GenerateNumberRangeMessage(input)

		if genErr != nil {
			log.Println("Invalid input provided")
		}

		fmt.Println(message)
	}
}

func GenerateNumberRangeMessage(input string) (string, error) {
	inputNum, err := convertToNumber(input)
	var output string

	if err != nil {
		return output, err
	}

	if inputNum >= 1 && inputNum <= 10 {
		output = fmt.Sprintln(inputNum, "is between 1 and 10")
	} else {
		output = fmt.Sprintln(inputNum, "is not between 1 and 10")
	}

	return output, nil
}

func convertToNumber(input string) (int, error) {
	output, err := strconv.Atoi(input)

	return output, err
}
