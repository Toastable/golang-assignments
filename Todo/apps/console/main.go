package main

import (
	"bufio"
	"common"
	"fmt"
	"os"
	"strings"
	"todo_inmemory_service"
	"todo_service"
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
	case "list":
		handleGet(service)
	case "exit":
		handleExit()
	}
}

func handleGet(service *todo_inmemory_service.TodoService) {
	todos, err := service.GetAll()
	handleError(err)

	printTodosToConsole(todos)
}

func handleCreate(service *todo_inmemory_service.TodoService, inputArgs string) {
	text := strings.Split(inputArgs, " ")[1:]

	if len(text) == 0 {
		fmt.Println("Provide a description of your task")
	} else {
		newTodoText := strings.Join(strings.Split(inputArgs, " ")[1:], " ")
		err := service.Create(newTodoText[:50], false)
		handleError(err)
	}
}

func printTodosToConsole(todos []todo_service.Todo) {
	fmt.Println("|--------------------------------------|----------------------------------------------------|-------------|")
	fmt.Printf("| %-36s | %-50s | %-10s | \n", "ID", "Description", "Is Complete")
	fmt.Println("|--------------------------------------|----------------------------------------------------|-------------|")

	for _, todo := range todos {
		fmt.Println(todo.String())
	}

	fmt.Println("|--------------------------------------|----------------------------------------------------|-------------|")
}

func handleExit() {
	os.Exit(1)
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v \n", err)
	}
}
