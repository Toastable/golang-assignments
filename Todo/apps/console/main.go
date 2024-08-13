package main

import (
	"bufio"
	"common"
	"fmt"
	"os"
	"strings"
	"todo_inmemory_service"
)

var initialConfigFilePath = "todos.json"

func main() {
	todoService := todo_inmemory_service.TodoService{}

	common.PopulateInMemoryTodos(&todoService, initialConfigFilePath)

	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Todo List Application")
	fmt.Println("|-------------------------------------------------------|")
	fmt.Println("Enter your commands or enter help for a list of commands")

	for inputScanner.Scan() {
		inputArgs := inputScanner.Text()

		processCommand(&todoService, inputArgs)
	}
}

func processCommand(service *todo_inmemory_service.TodoService, inputArguments string) {
	inputCommand := strings.Split(inputArguments, " ")[0]

	switch inputCommand {
	case "create":
		handleCreate(service, inputArguments)
	case "exit":
		handleExit()
	}
}

func handleCreate(service *todo_inmemory_service.TodoService, inputArgs string) {
	text := strings.Split(inputArgs, " ")[1:]

	if len(text) == 0 {
		fmt.Println("Provide a description of your task")
	} else {
		newTodoText := strings.Join(strings.Split(inputArgs, " ")[1:], " ")

		err := service.Create(newTodoText)
		handleError(err)
	}
}

func handleExit() {
	os.Exit(1)
}

func handleError(err error) {
	fmt.Fprintf(os.Stdout, "%v \n", err)
}
