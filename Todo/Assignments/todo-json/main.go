package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"todo_inmemory_service"
	"todo_service"
)

const todoFilePath = "todos.json"

func main() {
	todos := populateInMemoryTodos()
	printTodosToConsole(todos)
}

func printTodosToConsole(todos []todo_service.Todo) {
	fmt.Println("|----------------------------------------------------|-------------|")
	fmt.Printf("| %-50s | %-10s | \n", "Description", "Is Complete")
	fmt.Println("|----------------------------------------------------|-------------|")

	for _, todo := range todos {
		fmt.Printf("| %-50s | %-11t | \n", todo.Text, todo.Status)
	}

	fmt.Println("|----------------------------------------------------|-------------|")
}

func populateInMemoryTodos() []todo_service.Todo {
	inMemoryTodoService := todo_inmemory_service.TodoService{}
	todos := readJsonFileFromDisk(todoFilePath)

	for _, todo := range todos {
		inMemoryTodoService.Create(todo.Text, todo.Status)
	}

	return todos
}

func readJsonFileFromDisk(filePath string) []todo_service.Todo {

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	var todos []todo_service.Todo
	err = json.Unmarshal(content, &todos)
	if err != nil {
		log.Fatal("Failed to unmarshal file contents: ", err)
	}

	return todos
}
