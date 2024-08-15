package main

import (
	"common"
	"net/http"
	"server"
	"todo_inmemory_service"
)

const address = ":3000"

var initialConfigFilePath = "todos.json"

func main() {
	httpServer := http.NewServeMux()
	todoService := todo_inmemory_service.TodoService{}
	common.PopulateInMemoryTodos(&todoService, initialConfigFilePath)

	setupRouting(httpServer, &todoService)

	http.ListenAndServe(address, httpServer)
}

func setupRouting(httpServer *http.ServeMux, service *todo_inmemory_service.TodoService) {
	httpServer.Handle("GET /api/todo/", http.HandlerFunc(server.GetAllHandler(service)))
	httpServer.Handle("POST /api/todo/", http.HandlerFunc(server.CreateHandler(service)))
}
