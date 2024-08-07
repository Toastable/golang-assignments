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

		inputNum := convertToNumber(input)

		if inputNum >= 1 && inputNum <= 10 {
			fmt.Fprintln(os.Stdout, inputNum, "is between 1 and 10")
		} else {
			fmt.Fprintln(os.Stdout, inputNum, "is not between 1 and 10")
		}
	}
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
