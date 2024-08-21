package main

import (
	"common"
	"net/http"
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
	httpServer.Handle("GET /api/todo", http.HandlerFunc(GetAllHandler(service)))
	httpServer.Handle("GET /api/todo/{id}", http.HandlerFunc(GetHandler(service)))
	httpServer.Handle("POST /api/todo", http.HandlerFunc(CreateHandler(service)))
	httpServer.Handle("PATCH /api/todo", http.HandlerFunc(UpdateHandler(service)))
	httpServer.Handle("DELETE /api/todo/{id}", http.HandlerFunc(DeleteHandler(service)))
	httpServer.Handle("GET /api/alive", http.HandlerFunc(AliveHandler))
}
