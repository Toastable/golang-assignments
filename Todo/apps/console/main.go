package main

import (
	"bufio"
	"common"
	"fmt"
	"os"
	"strconv"
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
	case "delete":
		handleDelete(service, inputArguments)
	case "update":
		handleUpdate(service, inputArguments)
	case "help":
		handleHelp()
	case "exit":
		handleExit()
	}
}

func handleHelp() {
	fmt.Println("The following commands are available: ")
	fmt.Println("create: accepts a task text and creates a new task using it")
	fmt.Println("list : returns a list of all currently held tasks, their IDs and their Status")
	fmt.Println("delete : accepts an ID and deletes the associated task it if it exists")
	fmt.Println("update : accepts an ID and a Status and updates the task if it exists")
	fmt.Println("exit : shuts down the programme")
}

func handleUpdate(service *todo_inmemory_service.TodoService, inputArguments string) {
	argsArray := strings.Split(inputArguments, " ")

	if len(argsArray) < 3 {
		fmt.Println("Provide an ID and Status of a task to delete")
	} else {
		newStatus, parseError := strconv.ParseBool(argsArray[2])

		if parseError != nil {
			handleError(parseError)
		}

		fmt.Println(newStatus)

		err := service.Update(argsArray[1], newStatus)

		if err != nil {
			handleError(err)
		} else {
			fmt.Println("Task successfully updated")
		}
	}
}

func handleDelete(service *todo_inmemory_service.TodoService, inputArguments string) {
	argsArray := strings.Split(inputArguments, " ")

	if len(argsArray) == 0 {
		fmt.Println("Provide an ID of a task to delete")
	} else {
		err := service.Delete(argsArray[1])

		if err != nil {
			handleError(err)
		} else {
			fmt.Println("Task successfully deleted")
		}
	}
}

func handleGet(service *todo_inmemory_service.TodoService) {
	todos, err := service.GetAll()

	if err != nil {
		handleError(err)
	} else {
		printTodosToConsole(todos)
	}
}

func handleCreate(service *todo_inmemory_service.TodoService, inputArgs string) {
	text := strings.Split(inputArgs, " ")[1:]

	if len(text) == 0 {
		fmt.Println("Provide a description of your task")
	} else {
		newTodoText := strings.Join(strings.Split(inputArgs, " ")[1:], " ")
		err := service.Create(newTodoText[:50], false)

		if err != nil {
			handleError(err)
		} else {
			fmt.Println("Task created successfully")
		}
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
	fmt.Fprintf(os.Stdout, "%v \n", err)
}
