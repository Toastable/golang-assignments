package main

import (
	"log"
	"os"
)

const todoFilePath = "todos.json"

func main() {

}

func populateInMemoryTodos() {
	//todos := readFileFromDisk(todoFilePath)

	// for _, todo := range todos {

	// }
}

func readFileFromDisk(filePath string) string {
	contentBytes, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(contentBytes)
}
