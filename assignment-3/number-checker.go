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
		handleError(err)

		fmt.Println(GenerateNumberRangeMessage(input))
	}
}

func GenerateNumberRangeMessage(input string) string {
	inputNum := convertToNumber(input)
	var output string

	if inputNum >= 1 && inputNum <= 10 {
		output = fmt.Sprintln(inputNum, "is between 1 and 10")
	} else {
		output = fmt.Sprintln(inputNum, "is not between 1 and 10")
	}

	return output
}

func convertToNumber(input string) int {
	output, err := strconv.Atoi(input)
	handleError(err)

	return output
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
