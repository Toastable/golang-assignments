package common

import (
	"encoding/json"
	"log"
	"os"
	"todo_inmemory_service"
	"todo_service"
)

func PopulateInMemoryTodos(service *todo_inmemory_service.TodoService, todoFilePath string) []todo_service.Todo {

	todos := readJsonFileFromDisk(todoFilePath)

	for _, todo := range todos {
		service.Create(todo.Text)
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
