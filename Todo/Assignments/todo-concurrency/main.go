package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"todo_inmemory_service"
	"todo_service"
)

const todoFilePath = "todos.json"

func main() {

	todoChannel := make(chan todo_service.Todo)
	statusChannel := make(chan bool)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	todos := populateInMemoryTodos()

	go printTextToConsole(&waitGroup, todoChannel, statusChannel, todos)
	go printStatusToConsole(&waitGroup, todoChannel, statusChannel, todos)

	waitGroup.Wait()
}

func printTextToConsole(waitGroup *sync.WaitGroup, todoChannel chan todo_service.Todo, statusChannel chan bool, todos []todo_service.Todo) {
	defer waitGroup.Done()

	for _, todo := range todos {
		fmt.Printf("Task: %s \n", todo.Text)

		todoChannel <- todo
		<-statusChannel
	}
}

func printStatusToConsole(waitGroup *sync.WaitGroup, todoChannel chan todo_service.Todo, statusChannel chan bool, todos []todo_service.Todo) {
	defer waitGroup.Done()

	for _, todo := range todos {
		<-todoChannel
		fmt.Printf("Is Complete: %t \n", todo.Status)
		statusChannel <- todo.Status
	}
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
