package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter your Details")
	fmt.Println("|----------------------|")

	var firstName, middleName, lastName string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your First Name")
	firstName = readFromConsole(*scanner)

	fmt.Println("Enter your Middle Names")
	middleName = readFromConsole(*scanner)

	fmt.Println("Enter your Last Name")
	lastName = readFromConsole(*scanner)

	fmt.Printf(GenerateFullNameMessage(firstName, middleName, lastName))
}

func GenerateFullNameMessage(firstName, middleName, lastName string) string {
	return fmt.Sprintf("You entered %s %s %s", firstName, middleName, lastName)
}

func readFromConsole(scanner bufio.Scanner) string {
	scanner.Scan()
	err := scanner.Err()
	handleError(err)

	return scanner.Text()
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
